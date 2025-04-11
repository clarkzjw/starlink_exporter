<p align="center">
  <h3 align="center">Starlink Prometheus Exporter Monitoring Stack</h3>
</p>

---

A [Starlink](https://www.starlink.com/) exporter for Prometheus. Not affiliated with or acting on behalf of Starlink(™)

[![build](https://github.com/clarkzjw/starlink_exporter/actions/workflows/build.yaml/badge.svg)](https://github.com/clarkzjw/starlink_exporter/actions/workflows/build.yaml)
[![License](https://img.shields.io/github/license/clarkzjw/starlink_exporter)](/LICENSE)
[![Release](https://img.shields.io/github/release/clarkzjw/starlink_exporter.svg)](https://github.com/clarkzjw/starlink_exporter/releases/latest)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/clarkzjw/starlink_exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/clarkzjw/starlink_exporter)](https://goreportcard.com/report/github.com/clarkzjw/starlink_exporter)

Original repository:

+ https://github.com/danopstech/starlink
+ https://github.com/danopstech/starlink_exporter

---

Last updated with firmware version: `2025.04.08.cr53207 / 05de8289-7bcc-476b-ad62-8cf8cc2a73fe.uterm_manifest.release`

Starlink gRPC protobuf for Golang: [clarkzjw/starlink-grpc-golang](https://github.com/clarkzjw/starlink-grpc-golang)

Starlink dish firmware tracking website: https://starlinktrack.com/firmware/dishy

---

## Usage:

### Flags

`starlink_exporter` is configured through the use of optional command line flags

```bash
$ ./starlink_exporter -h
Usage of ./starlink_exporter:
  -address string
        IP address and port to reach dish (default "192.168.100.1:9200")
  -port string
        listening port to expose metrics on (default "9817")
```

### Binaries

For pre-built binaries please take a look at the [releases](https://github.com/clarkzjw/starlink_exporter/releases).

```bash
./starlink_exporter [flags]
```

### Docker Compose Stack

```bash
docker-compose up -d
```

## Grafana Dashboard:

<p align="center">
	<img src="https://github.com/clarkzjw/starlink_exporter/raw/main/static/screenshot.png" width="95%">
</p>
