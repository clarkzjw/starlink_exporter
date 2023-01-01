package exporter

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/clarkzjw/starlink_exporter/pkg/spacex.com/api/device"
)

const (
	// DishAddress to reach Starlink dish ip:port
	DishAddress   = "192.168.100.1:9200"
	RouterAddress = "192.168.1.1:9000"
	namespace     = "starlink"
)

var (
	// RouterInfo
	routerInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "router", "info"),
		"Running router software versions and IDs",
		[]string{
			"id",
			"hardware_version",
			"software_version",
			"manufactured_version",
			"country_code",
			"utc_offset_s",
			"software_partitions_equal",
			"is_dev",
			"boot_count",
			"anti_rollback_version",
			"is_hitl",
			"last_boot_reason",
		}, nil,
	)

	// Router IPv4 Wan Address
	routerIPv4WanAddress = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "router", "ipv4_wan_address"),
		"IPv4 WAN Address", []string{"ipv4_wan_address"}, nil)

	// Several ping latency metrics on the router
	routerToDishPingLatencyMs = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "router", "router_to_dish_ping_latency_ms"),
		"Router to Dish Ping latency ms", nil, nil,
	)

	routerPopPingLatencyMs = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "router", "router_pop_ping_latency_ms"),
		"Router Pop Ping latency ms", nil, nil,
	)

	routerPingLatencyMs = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "router", "router_ping_latency_ms"),
		"Router Ping latency ms", nil, nil,
	)

	ethernetInterface = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "router", "ethernet_interface"),
		"Ethernet interface stats on the router",
		[]string{
			"name",
			"rx_bytes",
			"rx_packets",
			"tx_bytes",
			"tx_packets",
			"duplex",
		}, nil,
	)

	// Router Uptime Seconds
	routerUptimeS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "router", "uptime_s"),
		"Router uptime seconds", nil, nil,
	)

	// DeviceInfo
	dishInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "info"),
		"Running software versions and IDs of hardware",
		[]string{
			"device_id",
			"hardware_version",
			"software_version",
			"manufactured_version",
			"country_code",
			"bootcount",
			"utc_offset"}, nil,
	)
	dishConfig = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "dish_config"),
		"Dish Config",
		[]string{
			"snow_melt_mode",
			"location_request_mode",
			"level_dish_mode",
			"power_save_mode",
		}, nil,
	)
	SoftwarePartitionsEqual = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "software_partitions_equal"),
		"Starlink Dish Software Partitions Equal.",
		nil, nil,
	)

	dishMobilityClass = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "mobility_class"),
		"Dish mobility class",
		[]string{"mobility_class"}, nil,
	)

	userClassOfService = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "class_of_service"),
		"User class of service",
		[]string{"class_of_service"}, nil,
	)

	dishReadyState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "ready_state"),
		"Dish ready states",
		[]string{
			"cady",
			"scp",
			"l1l2",
			"xphy",
			"aap",
			"rf",
		}, nil,
	)

	dishIsDev = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "is_dev"),
		"Starlink Dish is Dev.",
		nil, nil,
	)
	dishBootCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "boot_count"),
		"Starlink Dish boot count.",
		nil, nil,
	)
	dishAntiRollbackVersion = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "anti_rollback_version"),
		"Starlink Dish Anti Rollback Version.",
		nil, nil,
	)
	dishIsHit = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "is_hit"),
		"Starlink Dish is Hit.",
		nil, nil,
	)

	// BootInfo
	dishBootInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "info_debug"),
		"Debug Dish Info",
		[]string{
			"count_by_reason",
			"count_by_reason_delta",
			"last_reason",
			"last_count"}, nil,
	)

	// DeviceState
	dishUp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "up"),
		"Was the last query of Starlink dish successful.",
		nil, nil,
	)
	dishScrapeDurationSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "scrape_duration_seconds"),
		"Time to scrape metrics from starlink dish",
		nil, nil,
	)
	dishUptimeSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "uptime_seconds"),
		"Dish running time",
		nil, nil,
	)

	// DishOutages
	dishOutage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "outage_duration"),
		"Starlink Dish Outage Information",
		[]string{"start_time", "cause"}, nil,
	)
	dishOutageDidSwitch = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "outage_did_switch"),
		"Starlink Dish Outage Information",
		nil, nil,
	)

	// DishGpsStats
	dishGpsValid = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "gps_valid"),
		"GPS Status.",
		nil, nil,
	)
	dishGpsSats = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "gps_sats"),
		"Number of GPS Sats.",
		nil, nil,
	)

	// DishStatus
	dishSecondsToFirstNonemptySlot = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "first_nonempty_slot_seconds"),
		"Seconds to next non empty slot",
		nil, nil,
	)
	dishPopPingDropRatio = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "pop_ping_drop_ratio"),
		"Percent of pings dropped",
		nil, nil,
	)
	dishDownlinkThroughputBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "downlink_throughput_bytes"),
		"Amount of bandwidth in bytes per second download",
		nil, nil,
	)
	dishUplinkThroughputBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "uplink_throughput_bytes"),
		"Amount of bandwidth in bytes per second upload",
		nil, nil,
	)
	dishPopPingLatencySeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "pop_ping_latency_seconds"),
		"Latency of connection in seconds",
		nil, nil,
	)
	dishStowRequested = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "dish_stow_requested"),
		"stow requested",
		nil, nil,
	)
	dishBoreSightAzimuthDeg = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "bore_sight_azimuth_deg"),
		"azimuth in degrees",
		nil, nil,
	)
	dishBoreSightElevationDeg = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "bore_sight_elevation_deg"),
		"elevation in degrees",
		nil, nil,
	)
	dishEthSpeedMbps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "eth_speed"),
		"ethernet speed",
		nil, nil,
	)

	// DishAlerts
	dishAlertRoaming = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_roaming"),
		"Status of roaming",
		nil, nil,
	)
	dishAlertMotorsStuck = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_motors_stuck"),
		"Status of motor stuck",
		nil, nil,
	)
	dishAlertThermalThrottle = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_thermal_throttle"),
		"Status of thermal throttling",
		nil, nil,
	)
	dishAlertThermalShutdown = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_thermal_shutdown"),
		"Status of thermal shutdown",
		nil, nil,
	)
	dishAlertMastNotNearVertical = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_mast_not_near_vertical"),
		"Status of mast position",
		nil, nil,
	)
	dishUnexpectedLocation = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_unexpected_location"),
		"Status of location",
		nil, nil,
	)
	dishSlowEthernetSpeeds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_slow_eth_speeds"),
		"Status of ethernet",
		nil, nil,
	)
	dishInstallPending = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_install_pending"),
		"Installation Pending",
		nil, nil,
	)
	dishIsHeating = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_is_heating"),
		"Is Heating",
		nil, nil,
	)

	// DishObstructions
	dishCurrentlyObstructed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "currently_obstructed"),
		"Status of view of the sky",
		nil, nil,
	)
	dishFractionObstructionRatio = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "fraction_obstruction_ratio"),
		"Percentage of obstruction",
		nil, nil,
	)
	dishValidSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "valid_seconds"),
		"Unknown",
		nil, nil,
	)
	dishProlongedObstructionDurationSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "prolonged_obstruction_duration_seconds"),
		"Average in seconds of prolonged obstructions",
		nil, nil,
	)
	dishProlongedObstructionIntervalSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "prolonged_obstruction_interval_seconds"),
		"Average prolonged obstruction interval in seconds",
		nil, nil,
	)
	dishProlongedObstructionValid = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "prolonged_obstruction_valid"),
		"Average prolonged obstruction is valid",
		nil, nil,
	)
	dishWedgeFractionObstructionRatio = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "wedge_fraction_obstruction_ratio"),
		"Percentage of obstruction per wedge section",
		[]string{"wedge", "wedge_name"}, nil,
	)
	dishWedgeAbsFractionObstructionRatio = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "wedge_abs_fraction_obstruction_ratio"),
		"Percentage of Absolute fraction per wedge section",
		[]string{"wedge", "wedge_name"}, nil,
	)
)

// Exporter collects Starlink stats from the Dish and exports them using
// the prometheus metrics package.
type Exporter struct {
	Conn         *grpc.ClientConn
	RouterConn   *grpc.ClientConn
	Client       device.DeviceClient
	RouterClient device.DeviceClient

	DishID      string
	RouterID    string
	CountryCode string
}

// New returns an initialized Exporter.
func New(address string) (*Exporter, error) {
	ctx, connCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer connCancel()
	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("error creating underlying gRPC connection to starlink dish: %s", err.Error())
	}

	routerCtx, routerConnCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer routerConnCancel()
	routerConn, err := grpc.DialContext(routerCtx, RouterAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("error creating underlying gRPC connection to starlink router: %s", err.Error())
	}

	ctx, HandleCancel := context.WithTimeout(context.Background(), time.Second*1)
	defer HandleCancel()

	resp, err := device.NewDeviceClient(conn).Handle(ctx, &device.Request{
		Request: &device.Request_GetDeviceInfo{},
	})

	if err != nil {
		return nil, fmt.Errorf("could not collect inital information from dish: %s", err.Error())
	}

	routerCtx, routerHandleCancel := context.WithTimeout(context.Background(), time.Second*1)
	defer routerHandleCancel()
	routerResp, err := device.NewDeviceClient(routerConn).Handle(routerCtx, &device.Request{
		Request: &device.Request_GetDeviceInfo{},
	})
	if err != nil {
		return nil, fmt.Errorf("could not collect inital information from router: %s", err.Error())
	}

	return &Exporter{
		Conn:         conn,
		RouterConn:   routerConn,
		Client:       device.NewDeviceClient(conn),
		RouterClient: device.NewDeviceClient(routerConn),
		DishID:       resp.GetGetDeviceInfo().GetDeviceInfo().GetId(),
		RouterID:     routerResp.GetGetDeviceInfo().GetDeviceInfo().GetId(),
		CountryCode:  resp.GetGetDeviceInfo().GetDeviceInfo().GetCountryCode(),
	}, nil
}

// Describe describes all the metrics ever exported by the Starlink exporter. It
// implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	// Router
	ch <- routerInfo
	ch <- routerIPv4WanAddress
	ch <- routerToDishPingLatencyMs
	ch <- routerPopPingLatencyMs
	ch <- routerPingLatencyMs
	ch <- routerUptimeS

	ch <- ethernetInterface

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
	ch <- dishScrapeDurationSeconds

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
}

// Collect fetches the stats from Starlink dish and delivers them
// as Prometheus metrics. It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	start := time.Now()

	ok := e.collectDishStatus(ch)
	ok = ok && e.collectRouterStatus(ch)
	ok = ok && e.collectDishObstructions(ch)
	ok = ok && e.collectDishAlerts(ch)
	ok = ok && e.collectDishConfig(ch)
	ok = ok && e.collectNetworkInterface(ch)

	if ok {
		ch <- prometheus.MustNewConstMetric(
			dishUp, prometheus.GaugeValue, 1.0,
		)
		ch <- prometheus.MustNewConstMetric(
			dishScrapeDurationSeconds, prometheus.GaugeValue, time.Since(start).Seconds(),
		)
	} else {
		ch <- prometheus.MustNewConstMetric(
			dishUp, prometheus.GaugeValue, 0.0,
		)
	}
}

// TODO
func (e *Exporter) collectNetworkInterface(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetNetworkInterfaces{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	resp, err := e.RouterClient.Handle(ctx, req)
	if err != nil {
		log.Errorf("failed to get dish config from dish: %s", err.Error())
		return false
	}

	netIfaces := resp.GetGetNetworkInterfaces().GetNetworkInterfaces()
	for _, iface := range netIfaces {
		if iface.GetEthernet() != nil {
			ch <- prometheus.MustNewConstMetric(
				ethernetInterface, prometheus.GaugeValue, 1.00,
				iface.GetName(),
				fmt.Sprint(iface.GetRxStats().GetBytes()),
				fmt.Sprint(iface.GetRxStats().GetPackets()),
				fmt.Sprint(iface.GetTxStats().GetBytes()),
				fmt.Sprint(iface.GetTxStats().GetPackets()),
				iface.GetEthernet().GetDuplex().String(),
			)
		}
	}
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
		log.Errorf("failed to get dish config from dish: %s", err.Error())
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

func (e *Exporter) collectRouterStatus(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetStatus{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := e.RouterClient.Handle(ctx, req)
	if err != nil {
		log.Errorf("failed to collect status from router: %s", err.Error())
		return false
	}

	routerStatus := resp.GetWifiGetStatus()
	routerI := routerStatus.GetDeviceInfo()

	ch <- prometheus.MustNewConstMetric(
		routerInfo, prometheus.GaugeValue, 1.00,
		routerI.GetId(),
		routerI.GetHardwareVersion(),
		routerI.GetSoftwareVersion(),
		routerI.GetManufacturedVersion(),
		routerI.GetCountryCode(),
		fmt.Sprint(routerI.GetUtcOffsetS()),
		fmt.Sprint(routerI.GetSoftwarePartitionsEqual()),
		fmt.Sprint(routerI.GetIsDev()),
		fmt.Sprint(routerI.GetBootcount()),
		fmt.Sprint(routerI.GetAntiRollbackVersion()),
		fmt.Sprint(routerI.GetIsHitl()),
		routerI.GetBoot().GetLastReason().String(),
	)

	ch <- prometheus.MustNewConstMetric(
		routerIPv4WanAddress, prometheus.GaugeValue, 1.00,
		routerStatus.GetIpv4WanAddress())

	ch <- prometheus.MustNewConstMetric(
		routerUptimeS, prometheus.CounterValue, float64(routerStatus.GetDeviceState().GetUptimeS()))

	ch <- prometheus.MustNewConstMetric(
		routerToDishPingLatencyMs, prometheus.GaugeValue, float64(routerStatus.GetDishPingLatencyMs()))

	ch <- prometheus.MustNewConstMetric(
		routerPopPingLatencyMs, prometheus.GaugeValue, float64(routerStatus.GetPopPingLatencyMs()))

	ch <- prometheus.MustNewConstMetric(
		routerPingLatencyMs, prometheus.GaugeValue, float64(routerStatus.GetPingLatencyMs()))

	return true
}

func (e *Exporter) collectDishStatus(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetStatus{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("failed to collect status from dish: %s", err.Error())
		return false
	}

	dishStatus := resp.GetDishGetStatus()
	dishI := dishStatus.GetDeviceInfo()
	dishB := dishI.GetBoot()
	dishS := dishStatus.GetDeviceState()
	dishG := dishStatus.GetGpsStats()
	dishO := dishStatus.GetOutage()
	dishR := dishStatus.GetReadyStates()

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

	// DeviceInfo
	ch <- prometheus.MustNewConstMetric(
		dishInfo, prometheus.GaugeValue, 1.00,
		dishI.GetId(),
		dishI.GetHardwareVersion(),
		dishI.GetSoftwareVersion(),
		dishI.GetManufacturedVersion(),
		dishI.GetCountryCode(),
		fmt.Sprint(dishI.GetBootcount()),
		fmt.Sprint(dishI.GetUtcOffsetS()),
	)
	ch <- prometheus.MustNewConstMetric(
		SoftwarePartitionsEqual, prometheus.GaugeValue, flool(dishI.GetSoftwarePartitionsEqual()),
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

	// BootInfo - Need to expand this!
	ch <- prometheus.MustNewConstMetric(
		dishBootInfo, prometheus.GaugeValue, 1.00,
		fmt.Sprint(dishB.GetCountByReason()),
		fmt.Sprint(dishB.GetCountByReasonDelta()),
		fmt.Sprint(dishB.GetLastReason()),
		fmt.Sprint(dishB.GetLastCount()),
	)

	// DeviceState
	ch <- prometheus.MustNewConstMetric(
		dishUptimeSeconds, prometheus.CounterValue, float64(dishS.GetUptimeS()),
	)

	// DishOutage
	ch <- prometheus.MustNewConstMetric(
		dishOutage, prometheus.GaugeValue, float64(dishO.GetDurationNs()),
		fmt.Sprint(dishO.GetStartTimestampNs()),
		dishO.GetCause().String(),
	)
	ch <- prometheus.MustNewConstMetric(
		dishOutageDidSwitch, prometheus.GaugeValue, flool(dishO.GetDidSwitch()),
	)

	// DishGpsStats
	ch <- prometheus.MustNewConstMetric(
		dishGpsValid, prometheus.GaugeValue, flool(dishG.GetGpsValid()),
	)
	ch <- prometheus.MustNewConstMetric(
		dishGpsSats, prometheus.GaugeValue, float64(dishG.GetGpsSats()),
	)

	// DishStatus
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

	return true
}

func (e *Exporter) collectDishObstructions(ch chan<- prometheus.Metric) bool {
	req := &device.Request{
		Request: &device.Request_GetStatus{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	resp, err := e.Client.Handle(ctx, req)
	if err != nil {
		log.Errorf("failed to collect obstructions from dish: %s", err.Error())
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
		dishValidSeconds, prometheus.CounterValue, float64(obstructions.GetValidS()),
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

	//for i, v := range obstructions.GetWedgeFractionObstructed() {
	//	ch <- prometheus.MustNewConstMetric(
	//		dishWedgeFractionObstructionRatio, prometheus.GaugeValue, float64(v),
	//		strconv.Itoa(i),
	//		fmt.Sprintf("%d_to_%d", i*30, (i+1)*30),
	//	)
	//}
	//
	//for i, v := range obstructions.GetWedgeAbsFractionObstructed() {
	//	ch <- prometheus.MustNewConstMetric(
	//		dishWedgeAbsFractionObstructionRatio, prometheus.GaugeValue, float64(v),
	//		strconv.Itoa(i),
	//		fmt.Sprintf("%d_to_%d", i*30, (i+1)*30),
	//	)
	//}

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
		log.Errorf("failed to collect alerts from dish: %s", err.Error())
		return false
	}
	alerts := resp.GetDishGetStatus().GetAlerts()

	ch <- prometheus.MustNewConstMetric(
		dishAlertMotorsStuck, prometheus.GaugeValue, flool(alerts.GetMotorsStuck()),
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

func flool(b bool) float64 {
	if b {
		return 1.00
	}
	return 0.00
}
