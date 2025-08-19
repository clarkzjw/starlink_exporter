package publicip

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	namespace = "starlink"
)

var (
	// public IP and PoP info
	dishPublicIPPoP = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "public_ip_pop"),
		"Public IPv4 address of the Starlink dish",
		[]string{"public_ipv4", "public_ipv6", "pop_code_ipv4", "pop_code_ipv6"}, nil,
	)
	iface       = os.Getenv("IFACE")
	curlTimeout = "5"
)

type Exporter struct {
}

func New() (*Exporter, error) {
	if iface == "" {
		return nil, fmt.Errorf("IFACE Starlink network interface is not set")
	}
	return &Exporter{}, nil
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- dishPublicIPPoP
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	_ = e.collectPublicIP(ch)
}

func (e *Exporter) collectPublicIP(ch chan<- prometheus.Metric) bool {
	publicIPv4, err := curlGetPublicIP(4)
	if err != nil {
		log.Warnf("Failed to fetch public IPv4: %s", err.Error())
		publicIPv4 = ""
	}
	publicIPv6, err := curlGetPublicIP(6)
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

// func getPublicIP(ipVersion int) (string, error) {
// 	url := "https://ifconfig.io/all.json"
// 	version := "tcp4"
// 	if ipVersion == 6 {
// 		version = "tcp6"
// 	}
// 	myDialer := net.Dialer{}
// 	transport := http.DefaultTransport.(*http.Transport).Clone()
// 	transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
// 		return myDialer.DialContext(ctx, version, addr)
// 	}

// 	client := http.Client{
// 		Transport: transport,
// 	}

// 	resp, err := client.Get(url)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to fetch public IP: %s", err.Error())
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return "", fmt.Errorf("failed to fetch public IP, status code: %d", resp.StatusCode)
// 	}

// 	type ifconfigResp struct {
// 		IP string `json:"ip"`
// 	}

// 	var data ifconfigResp
// 	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
// 		return "", fmt.Errorf("failed to decode public IP response: %s", err.Error())
// 	}
// 	if data.IP == "" {
// 		return "", fmt.Errorf("public IP not found in response")
// 	}

// 	return data.IP, nil
// }

func curlGetPublicIP(ipVersion int) (string, error) {
	cmd := exec.Command("curl", "--connect-timeout", curlTimeout, "-s", fmt.Sprintf("-%d", ipVersion), "--interface", iface, "https://ifconfig.io")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute curl command: %s", err.Error())
	}

	ip := strings.TrimSpace(string(output))
	if net.ParseIP(ip) == nil {
		return "", fmt.Errorf("invalid IP address returned by curl: %s", ip)
	}

	return ip, nil
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
