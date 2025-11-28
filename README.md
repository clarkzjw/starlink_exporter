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

Use the `docker-compose.yaml` in the [contrib](./contrib) directory.

```bash
docker-compose up -d
```

If you want to build locally, you can use the following command instead

```bash
docker-compose up -d --build
```

Podman is also supported if user or system socket is enabled:

```bash
systemctl --user enable --now podman.socket
DOCKER_HOST=unix:///run/user/$(id -u)/podman/podman.sock docker-compose up -d # --build is optional
```

## Grafana Dashboard:

<p align="center">
	<img src="https://github.com/clarkzjw/starlink_exporter/raw/main/static/Screenshot.jpeg" width="95%">
</p>
