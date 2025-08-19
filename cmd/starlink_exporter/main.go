package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/clarkzjw/starlink_exporter/internal/exporter"
	publicip "github.com/clarkzjw/starlink_exporter/internal/publicIP"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

const (
	metricsPath           = "/metrics"
	infrequentMetricsPath = "/infrequentMetrics"
	healthPath            = "/health"
)

func main() {
	port := flag.String("port", "9817", "listening port to expose metrics on")
	address := flag.String("address", exporter.DishAddress, "IP address and port to reach dish")
	flag.Parse()

	if os.Getenv("STARLINK_GRPC_ADDR_PORT") != "" {
		*address = os.Getenv("STARLINK_GRPC_ADDR_PORT")
	}

	var exporterClient *exporter.Exporter
	var publicIPClient *publicip.Exporter
	var err error
	retryDelay := 1

	for {
		exporterClient, err = exporter.New(*address)
		if err == nil {
			break
		}

		log.Warnf("Failed to connect to Starlink dish: %s, retrying in %d seconds...", err.Error(), retryDelay)
		time.Sleep(time.Duration(retryDelay) * time.Second)

		publicIPClient, _ = publicip.New()
	}

	defer func() {
		if err := exporterClient.Conn.Close(); err != nil {
			log.Errorf("Failed to close exporter connection: %s", err.Error())
		}
	}()
	log.Infof("Dish id: %s", exporterClient.DishID)

	r := prometheus.NewRegistry()
	r.MustRegister(exporterClient)

	r1 := prometheus.NewRegistry()
	r1.MustRegister(publicIPClient)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<html>
             <head><title>Starlink Exporter</title></head>
             <body>
             <h1>Starlink Exporter</h1>
             <p><a href='` + metricsPath + `'>Metrics</a></p>
             <p><a href='/health'>Health (gRPC connection state to Starlink dish)</a></p>
             </body>
             </html>`))
	})

	http.HandleFunc(healthPath, func(w http.ResponseWriter, r *http.Request) {
		switch exporterClient.Conn.GetState() {
		case 0, 2:
			// Idle or Ready
			w.WriteHeader(http.StatusOK)
		case 1, 3:
			// Connecting or TransientFailure
			w.WriteHeader(http.StatusServiceUnavailable)
		case 4:
			// Shutdown
			w.WriteHeader(http.StatusInternalServerError)
		}
		_, _ = fmt.Fprintf(w, "%s\n", exporterClient.Conn.GetState())
	})

	http.Handle(metricsPath, promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	http.Handle(infrequentMetricsPath, promhttp.HandlerFor(r1, promhttp.HandlerOpts{}))

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
