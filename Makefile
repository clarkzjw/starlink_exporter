all:
	go build -o starlink_exporter cmd/starlink_exporter/main.go

clean:
	rm starlink_exporter
