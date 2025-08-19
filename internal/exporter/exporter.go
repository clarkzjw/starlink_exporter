package exporter

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"net"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	device "github.com/clarkzjw/starlink-grpc-golang/pkg/spacex.com/api/device"
)

// Exporter collects Starlink stats from the Dish and exports them using
// the prometheus metrics package.
type Exporter struct {
	Conn   *grpc.ClientConn
	Client device.DeviceClient

	DishID      string
	CountryCode string
}

// New returns an initialized Exporter.
func New(address string) (*Exporter, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connect to Starlink dish gRPC interface failed: %s", err.Error())
	}

	defer func() {
		if err != nil {
			if err = conn.Close(); err != nil {
				log.Errorf("Failed to close gRPC client: %s", err.Error())
			}
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	client := device.NewDeviceClient(conn)
	resp, err := client.Handle(ctx, &device.Request{
		Request: &device.Request_GetDeviceInfo{},
	})
	if err != nil {
		return nil, fmt.Errorf("gRPC GetDeviceInfo failed: %s", err.Error())
	}

	deviceInfo := resp.GetGetDeviceInfo().GetDeviceInfo()
	if deviceInfo == nil {
		return nil, fmt.Errorf("gRPC GetDeviceInfo failed: deviceInfo is nil")
	}

	return &Exporter{
		Conn:        conn,
		Client:      client,
		DishID:      deviceInfo.GetId(),
		CountryCode: deviceInfo.GetCountryCode(),
	}, nil
}

// Describe describes all the metrics ever exported by the Starlink exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	// WiFi
	ch <- dishConfig

	// DeviceInfo
	ch <- dishMobilityClass
	ch <- userClassOfService
	ch <- dishReadyState
	ch <- dishInfo
	ch <- SoftwarePartitionsEqual
	ch <- dishIsDev
	ch <- dishBootCount
	ch <- dishAntiRollbackVersion
	ch <- dishIsHit

	// BootInfo
	ch <- dishBootInfo

	// DeviceState
	ch <- dishUp
	ch <- dishUptimeSeconds
	// ch <- dishScrapeDurationSeconds

	// DishOutage
	ch <- dishOutage
	ch <- dishOutageDidSwitch

	// DishGpsStats
	ch <- dishGpsValid
	ch <- dishGpsSats

	// DishStatus
	ch <- dishSecondsToFirstNonemptySlot
	ch <- dishPopPingDropRatio
	ch <- dishDownlinkThroughputBytes
	ch <- dishUplinkThroughputBytes
	ch <- dishPopPingLatencySeconds
	ch <- dishStowRequested
	ch <- dishBoreSightAzimuthDeg
	ch <- dishBoreSightElevationDeg
	ch <- dishEthSpeedMbps

	// DishAlerts
	ch <- dishAlertRoaming
	ch <- dishAlertMotorsStuck
	ch <- dishAlertThermalThrottle
	ch <- dishAlertThermalShutdown
	ch <- dishAlertMastNotNearVertical
	ch <- dishUnexpectedLocation
	ch <- dishSlowEthernetSpeeds
	ch <- dishInstallPending
	ch <- dishIsHeating

	// DishObstructions
	ch <- dishCurrentlyObstructed
	ch <- dishFractionObstructionRatio
	ch <- dishValidSeconds
	ch <- dishProlongedObstructionDurationSeconds
	ch <- dishProlongedObstructionIntervalSeconds
	ch <- dishProlongedObstructionValid
	ch <- dishWedgeFractionObstructionRatio
	ch <- dishWedgeAbsFractionObstructionRatio

	// Public IP and PoP code
	ch <- dishPublicIPPoP
}

// Collect fetches the stats from Starlink dish and delivers them as Prometheus metrics.
// It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	ok := e.collectDishStatus(ch)
	ok = ok && e.collectDishLocation(ch)
	ok = ok && e.collectDishObstructionStatus(ch)
	ok = ok && e.collectDishObstructionMap(ch)
	ok = ok && e.collectDishAlerts(ch)
	ok = ok && e.collectDishConfig(ch)
	ok = ok && e.collectAlignmentStats(ch)
	ok = ok && e.collectDishDiagnostics(ch)
	ok = ok && e.collectDishPower(ch)
	ok = ok && e.collectPublicIP(ch)

	if ok {
		ch <- prometheus.MustNewConstMetric(
			dishUp, prometheus.GaugeValue, 1.0,
		)
	} else {
		ch <- prometheus.MustNewConstMetric(
			dishUp, prometheus.GaugeValue, 0.0,
		)
	}
}

func (e *Exporter) collectDishStatus(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetStatus{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("gRPC GetStatus failed: %s", err.Error())
		return false
	}

	dishStatus := resp.GetDishGetStatus()
	dishI := dishStatus.GetDeviceInfo()
	dishB := dishI.GetBoot()
	dishS := dishStatus.GetDeviceState()
	dishG := dishStatus.GetGpsStats()
	dishO := dishStatus.GetOutage()
	dishR := dishStatus.GetReadyStates()
	dishInit := dishStatus.GetInitializationDurationSeconds()
	dishQuaternion := dishStatus.GetNed2DishQuaternion()

	ch <- prometheus.MustNewConstMetric(
		dishNed2dishQuaternionQScalar, prometheus.GaugeValue, float64(dishQuaternion.QScalar))

	ch <- prometheus.MustNewConstMetric(
		dishNed2dishQuaternionQX, prometheus.GaugeValue, float64(dishQuaternion.QX))

	ch <- prometheus.MustNewConstMetric(
		dishNed2dishQuaternionQY, prometheus.GaugeValue, float64(dishQuaternion.QY))

	ch <- prometheus.MustNewConstMetric(
		dishNed2dishQuaternionQZ, prometheus.GaugeValue, float64(dishQuaternion.QZ))

	ch <- prometheus.MustNewConstMetric(
		dishInitializationDurationSeconds, prometheus.GaugeValue, 1.00,
		fmt.Sprint(dishInit.GetAttitudeInitialization()),
		fmt.Sprint(dishInit.GetBurstDetected()),
		fmt.Sprint(dishInit.GetEkfConverged()),
		fmt.Sprint(dishInit.GetFirstCplane()),
		fmt.Sprint(dishInit.GetFirstPopPing()),
		fmt.Sprint(dishInit.GetGpsValid()),
		fmt.Sprint(dishInit.GetInitialNetworkEntry()),
		fmt.Sprint(dishInit.GetNetworkSchedule()),
		fmt.Sprint(dishInit.GetRfReady()),
		fmt.Sprint(dishInit.GetStableConnection()),
	)

	ch <- prometheus.MustNewConstMetric(
		dishReadyState, prometheus.GaugeValue, 1.00,
		fmt.Sprint(dishR.GetCady()),
		fmt.Sprint(dishR.GetScp()),
		fmt.Sprint(dishR.GetL1L2()),
		fmt.Sprint(dishR.GetXphy()),
		fmt.Sprint(dishR.GetAap()),
		fmt.Sprint(dishR.GetRf()))

	ch <- prometheus.MustNewConstMetric(
		userClassOfService, prometheus.GaugeValue, 1.00,
		dishStatus.GetClassOfService().String())
	ch <- prometheus.MustNewConstMetric(
		dishMobilityClass, prometheus.GaugeValue, 1.00,
		dishStatus.GetMobilityClass().String())
	ch <- prometheus.MustNewConstMetric(
		dishInfo, prometheus.GaugeValue, 1.00,
		dishI.GetId(),
		dishI.GetBuildId(),
		dishI.GetHardwareVersion(),
		dishI.GetSoftwareVersion(),
		fmt.Sprint(dishI.GetGenerationNumber()),
		dishI.GetCountryCode(),
		fmt.Sprint(dishI.GetBootcount()),
		fmt.Sprint(dishI.GetUtcOffsetS()),
	)
	ch <- prometheus.MustNewConstMetric(
		SoftwarePartitionsEqual, prometheus.GaugeValue, flool(dishI.GetSoftwarePartitionsEqual()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishSoftwareUpdateState, prometheus.GaugeValue, 1.00, dishStatus.GetSoftwareUpdateState().String(),
	)
	ch <- prometheus.MustNewConstMetric(
		dishDisablementCode, prometheus.GaugeValue, 1.00, dishStatus.GetDisablementCode().String(),
	)
	ch <- prometheus.MustNewConstMetric(
		dishIsDev, prometheus.GaugeValue, flool(dishI.GetIsDev()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishBootCount, prometheus.CounterValue, float64(dishI.GetBootcount()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishAntiRollbackVersion, prometheus.CounterValue, float64(dishI.GetAntiRollbackVersion()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishIsHit, prometheus.GaugeValue, flool(dishI.GetIsHitl()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishBootInfo, prometheus.GaugeValue, 1.00,
		fmt.Sprint(dishB.GetCountByReason()),
		fmt.Sprint(dishB.GetCountByReasonDelta()),
		fmt.Sprint(dishB.GetLastReason()),
		fmt.Sprint(dishB.GetLastCount()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishUptimeSeconds, prometheus.CounterValue, float64(dishS.GetUptimeS()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishOutage, prometheus.GaugeValue, float64(dishO.GetDurationNs()),
		fmt.Sprint(dishO.GetStartTimestampNs()),
		dishO.GetCause().String(),
	)
	ch <- prometheus.MustNewConstMetric(
		dishOutageDidSwitch, prometheus.GaugeValue, flool(dishO.GetDidSwitch()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishGpsValid, prometheus.GaugeValue, flool(dishG.GetGpsValid()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishGpsSats, prometheus.GaugeValue, float64(dishG.GetGpsSats()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishSecondsToFirstNonemptySlot, prometheus.GaugeValue, float64(dishStatus.GetSecondsToFirstNonemptySlot()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishPopPingDropRatio, prometheus.GaugeValue, float64(dishStatus.GetPopPingDropRate()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishDownlinkThroughputBytes, prometheus.GaugeValue, float64(dishStatus.GetDownlinkThroughputBps()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishUplinkThroughputBytes, prometheus.GaugeValue, float64(dishStatus.GetUplinkThroughputBps()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishPopPingLatencySeconds, prometheus.GaugeValue, float64(dishStatus.GetPopPingLatencyMs()/1000),
	)
	ch <- prometheus.MustNewConstMetric(
		dishStowRequested, prometheus.GaugeValue, flool(dishStatus.GetStowRequested()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishBoreSightAzimuthDeg, prometheus.GaugeValue, float64(dishStatus.GetBoresightAzimuthDeg()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishBoreSightElevationDeg, prometheus.GaugeValue, float64(dishStatus.GetBoresightElevationDeg()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishEthSpeedMbps, prometheus.UntypedValue, float64(dishStatus.GetEthSpeedMbps()),
	)
	// ch <- prometheus.MustNewConstMetric(
	// 	dishPhyRxBeamSnrAvg, prometheus.GaugeValue, float64(dishStatus.GetPhyRxBeamSnrAvg()),
	// )
	// ch <- prometheus.MustNewConstMetric(
	// 	dishTemperateCenter, prometheus.GaugeValue, float64(dishStatus.GetTCenter()),
	// )
	return true
}

func getPublicIP(ipVersion int) (string, error) {
	url := "https://ifconfig.io/all.json"
	version := "tcp4"
	if ipVersion == 6 {
		version = "tcp6"
	}
	myDialer := net.Dialer{}
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return myDialer.DialContext(ctx, version, addr)
	}

	client := http.Client{
		Transport: transport,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch public IP: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch public IP, status code: %d", resp.StatusCode)
	}

	type ifconfigResp struct {
		IP string `json:"ip"`
	}

	var data ifconfigResp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("failed to decode public IP response: %s", err.Error())
	}
	if data.IP == "" {
		return "", fmt.Errorf("public IP not found in response")
	}

	return data.IP, nil
}

func getDnsPtrRecord(ip string) (string, error) {
	cmd := exec.Command("dig", "@1.1.1.1", "-x", ip, "+short")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute dig command: %s", err.Error())
	}

	ptrRecord := strings.TrimSpace(string(output))
	regex := `^customer\.(?P<pop>[a-z0-9]+)\.pop\.starlinkisp\.net\.$`
	re := regexp.MustCompile(regex)
	matches := re.FindStringSubmatch(ptrRecord)
	if len(matches) < 2 {
		return "", fmt.Errorf("failed to parse PTR record: %s", ptrRecord)
	}
	popCode := matches[1]
	if popCode == "" {
		return "", fmt.Errorf("pop code not found in PTR record: %s", ptrRecord)
	}

	return popCode, nil
}

func (e *Exporter) collectPublicIP(ch chan<- prometheus.Metric) bool {
	publicIPv4, err := getPublicIP(4)
	if err != nil {
		log.Warnf("Failed to fetch public IPv4: %s", err.Error())
		publicIPv4 = ""
	}
	publicIPv6, err := getPublicIP(6)
	if err != nil {
		log.Warnf("Failed to fetch public IPv6: %s", err.Error())
		publicIPv6 = ""
	}

	publicIPv4Ptr, err := getDnsPtrRecord(publicIPv4)
	if err != nil {
		log.Warnf("Failed to fetch public IPv4 PTR record: %s", err.Error())
		publicIPv4Ptr = ""
	}
	publicIPv6Ptr, err := getDnsPtrRecord(publicIPv6)
	if err != nil {
		log.Warnf("Failed to fetch public IPv6 PTR record: %s", err.Error())
		publicIPv6Ptr = ""
	}
	ch <- prometheus.MustNewConstMetric(
		dishPublicIPPoP, prometheus.GaugeValue, 1.00,
		fmt.Sprint(publicIPv4),
		fmt.Sprint(publicIPv6),
		fmt.Sprint(publicIPv4Ptr),
		fmt.Sprint(publicIPv6Ptr),
	)

	return true
}

func (e *Exporter) collectDishConfig(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_DishGetConfig{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("gRPC DishGetConfig failed: %s", err.Error())
		return false
	}

	dishC := resp.GetDishGetConfig()
	ch <- prometheus.MustNewConstMetric(
		dishConfig, prometheus.GaugeValue, 1.00,
		dishC.GetDishConfig().GetSnowMeltMode().String(),
		dishC.GetDishConfig().GetLocationRequestMode().String(),
		dishC.GetDishConfig().GetLevelDishMode().String(),
		fmt.Sprint(dishC.GetDishConfig().GetPowerSaveMode()),
	)
	return true
}

func (e *Exporter) collectDishLocation(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetLocation{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("gRPC GetLocation failed: %s", err.Error())
		// don't return false since location service might not be enabled
		return true
	}

	locationInfo := resp.GetGetLocation()

	locationSource := locationInfo.GetSource().String()
	sigmaM := locationInfo.GetSigmaM()
	horizontalSpeedMps := locationInfo.GetHorizontalSpeedMps()
	verticalSpeedMps := locationInfo.GetVerticalSpeedMps()

	lla := locationInfo.GetLla()
	lat := lla.GetLat()
	lon := lla.GetLon()
	alt := lla.GetAlt()

	ch <- prometheus.MustNewConstMetric(
		dishLocationInfo, prometheus.GaugeValue, 1.00,
		locationSource,
		fmt.Sprintf("%.6f", lat),
		fmt.Sprintf("%.6f", lon),
		fmt.Sprintf("%.3f", alt),
		fmt.Sprintf("%.6f", sigmaM),
		fmt.Sprintf("%.6f", horizontalSpeedMps),
		fmt.Sprintf("%.6f", verticalSpeedMps),
	)

	return true
}

func (e *Exporter) collectDishObstructionStatus(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetStatus{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("gRPC GetStatus failed: %s", err.Error())
		return false
	}

	obstructions := resp.GetDishGetStatus().GetObstructionStats()

	ch <- prometheus.MustNewConstMetric(
		dishCurrentlyObstructed, prometheus.GaugeValue, flool(obstructions.GetCurrentlyObstructed()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishFractionObstructionRatio, prometheus.GaugeValue, float64(obstructions.GetFractionObstructed()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishTimeObstructed, prometheus.GaugeValue, float64(obstructions.GetTimeObstructed()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishValidSeconds, prometheus.CounterValue, float64(obstructions.GetValidS()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishPatchesValid, prometheus.GaugeValue, float64(obstructions.GetPatchesValid()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishProlongedObstructionDurationSeconds, prometheus.GaugeValue, float64(obstructions.GetAvgProlongedObstructionDurationS()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishProlongedObstructionIntervalSeconds, prometheus.GaugeValue, float64(obstructions.GetAvgProlongedObstructionIntervalS()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishProlongedObstructionValid, prometheus.GaugeValue, flool(obstructions.GetAvgProlongedObstructionValid()),
	)

	return true
}

func (e *Exporter) collectDishObstructionMap(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_DishGetObstructionMap{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("gRPC GetStatus failed: %s", err.Error())
		return false
	}

	obstructionMap := resp.GetDishGetObstructionMap()

	rows := int(obstructionMap.NumRows)
	cols := int(obstructionMap.NumCols)
	referenceFrame := obstructionMap.GetMapReferenceFrame().String()
	data := obstructionMap.Snr

	upLeft := image.Point{0, 0}
	lowRight := image.Point{cols, rows}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < cols; x++ {
		for y := 0; y < rows; y++ {
			snr := data[y*cols+x]
			if snr > 1 {
				// shouldn't happen
				snr = 1.0
			}
			if snr == -1 {
				// background
				img.Set(x, y, color.Black)
			} else if snr > 0 {
				// use the same image color style as in starlink-grpc-tools
				// https://github.com/sparky8512/starlink-grpc-tools/blob/a3860e0a73d0b2280eed92eb8a2a97de0ea5fe43/dish_obstruction_map.py#L59-L87
				r := 255
				g := snr * 255
				b := snr * 255
				alpha := 255
				img.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(alpha)})
			}
		}
	}

	// Encode the image to PNG format in a buffer
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		fmt.Printf("Failed to encode image: %s", err.Error())
	}

	timestamp := time.Now().Format(time.RFC3339)
	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	ch <- prometheus.MustNewConstMetric(
		dishObstructionMap, prometheus.GaugeValue, 1.00,
		timestamp,
		fmt.Sprint(obstructionMap.GetNumRows()),
		fmt.Sprint(obstructionMap.GetNumCols()),
		fmt.Sprint(obstructionMap.GetMaxThetaDeg()),
		referenceFrame,
		fmt.Sprintf("data:image/png;base64,%s", b64),
	)

	return true
}

func (e *Exporter) collectDishDiagnostics(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetDiagnostics{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("gRPC GetStatus failed: %s", err.Error())
		return false
	}

	diagnostics := resp.GetDishGetDiagnostics()

	ch <- prometheus.MustNewConstMetric(
		dishGpsTimeS, prometheus.GaugeValue,
		float64(diagnostics.GetLocation().GetGpsTimeS()),
	)

	return true
}

func (e *Exporter) collectDishAlerts(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetStatus{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("gRPC GetStatus failed: %s", err.Error())
		return false
	}
	alerts := resp.GetDishGetStatus().GetAlerts()

	ch <- prometheus.MustNewConstMetric(
		dishAlertMotorsStuck, prometheus.GaugeValue, flool(alerts.GetMotorsStuck()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishPowerSupplyThermalThrottle, prometheus.GaugeValue, flool(alerts.GetPowerSupplyThermalThrottle()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishIsPowerSaveIdle, prometheus.GaugeValue, flool(alerts.GetIsPowerSaveIdle()),
	)
	// ch <- prometheus.MustNewConstMetric(
	// 	dishMovingWhileNotMobile, prometheus.GaugeValue, flool(alerts.GetMovingWhileNotMobile()),
	// )
	// ch <- prometheus.MustNewConstMetric(
	// 	dishMovingTooFastForPolicy, prometheus.GaugeValue, flool(alerts.GetMovingTooFastForPolicy()),
	// )
	ch <- prometheus.MustNewConstMetric(
		dishLowMotorCurrent, prometheus.GaugeValue, flool(alerts.GetLowMotorCurrent()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishLowerSignalThanPredicted, prometheus.GaugeValue, flool(alerts.GetLowerSignalThanPredicted()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishObstructionMapReset, prometheus.GaugeValue, flool(alerts.GetObstructionMapReset()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishAlertThermalThrottle, prometheus.GaugeValue, flool(alerts.GetThermalThrottle()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishAlertThermalShutdown, prometheus.GaugeValue, flool(alerts.GetThermalShutdown()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishAlertMastNotNearVertical, prometheus.GaugeValue, flool(alerts.GetMastNotNearVertical()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishUnexpectedLocation, prometheus.GaugeValue, flool(alerts.GetUnexpectedLocation()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishSlowEthernetSpeeds, prometheus.GaugeValue, flool(alerts.GetSlowEthernetSpeeds()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishAlertRoaming, prometheus.GaugeValue, flool(alerts.GetRoaming()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishInstallPending, prometheus.GaugeValue, flool(alerts.GetInstallPending()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishIsHeating, prometheus.GaugeValue, flool(alerts.GetIsHeating()),
	)

	return true
}

func (e *Exporter) collectAlignmentStats(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetStatus{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("gRPC GetStatus failed: %s", err.Error())
		return false
	}
	alignmentStats := resp.GetDishGetStatus().GetAlignmentStats()

	ch <- prometheus.MustNewConstMetric(
		dishAlignmentStats, prometheus.GaugeValue, 1.00,
		fmt.Sprint(alignmentStats.GetHasActuators()),
		fmt.Sprint(alignmentStats.GetActuatorState()),
		fmt.Sprint(alignmentStats.GetTiltAngleDeg()),
		fmt.Sprint(alignmentStats.GetBoresightAzimuthDeg()),
		fmt.Sprint(alignmentStats.GetBoresightElevationDeg()),
		fmt.Sprint(alignmentStats.GetAttitudeEstimationState()),
		fmt.Sprint(alignmentStats.GetAttitudeUncertaintyDeg()),
		fmt.Sprint(alignmentStats.GetDesiredBoresightAzimuthDeg()),
		fmt.Sprint(alignmentStats.GetDesiredBoresightElevationDeg()),
	)
	return true
}

func (e *Exporter) collectDishPower(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetHistory{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("gRPC GetStatus failed: %s", err.Error())
		return false
	}
	history := resp.GetDishGetHistory()
	powerHistory := history.GetPowerIn()

	latest_range, _, _ := computeSampleRange(history, 1)
	latest_index := latest_range[0]
	ch <- prometheus.MustNewConstMetric(
		dishPowerWatt, prometheus.GaugeValue, float64(powerHistory[latest_index]),
	)
	avg := 0.0
	for i := 0; i < len(powerHistory); i++ {
		avg += float64(powerHistory[i])
	}
	avg /= float64(len(powerHistory))
	ch <- prometheus.MustNewConstMetric(
		dishPowerWattAvg15min, prometheus.GaugeValue, avg,
	)
	return true
}

func flool(b bool) float64 {
	if b {
		return 1.00
	}
	return 0.00
}

// https://github.com/sparky8512/starlink-grpc-tools/blob/a3860e0a73d0b2280eed92eb8a2a97de0ea5fe43/starlink_grpc.py#L1038-L1090
func computeSampleRange(history *device.DishGetHistoryResponse, parseSamples int) ([]int, int, int) {
	current := int(history.Current)
	samples := len(history.PopPingDropRate)
	if samples == 0 {
		return []int{}, 0, 0
	}

	// Adjust parseSamples if needed
	if parseSamples < 0 || samples < parseSamples {
		parseSamples = samples
	}

	// Calculate start position
	start := current - parseSamples

	if start == current {
		return []int{}, 0, current
	}

	// Calculate ring buffer offsets
	endOffset := current % samples
	startOffset := start % samples

	// Create a slice to hold the range of sample indices
	var sampleRange []int

	// Set the range for the requested set of samples
	if startOffset < endOffset {
		// Continuous range
		for i := startOffset; i < endOffset; i++ {
			sampleRange = append(sampleRange, i)
		}
	} else {
		// Wrap-around range
		for i := startOffset; i < samples; i++ {
			sampleRange = append(sampleRange, i)
		}
		for i := 0; i < endOffset; i++ {
			sampleRange = append(sampleRange, i)
		}
	}

	return sampleRange, current - start, current
}
