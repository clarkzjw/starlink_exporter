# Starlink Exporter Monitoring System

This is a Starlink monitoring system based on Prometheus and Grafana. 

Forked from the original project by [danopstech/starlink](https://github.com/danopstech/starlink).

**Status**: 

* [The fork](https://github.com/clarkzjw/starlink_exporter) is currently updated to the `2d2db653-e245-403b-b1a5-1af2cca0aa43.uterm.release` firmware version.

* This documentation remains to be updated.

## User Guide

* Install

## Developer Guide

* Retrieve the latest protoset from the Starlink dish

[`get-protoset.sh`](../grpc/get-protoset.sh)

* gRPC Guide

Install tools:

```bash
sudo apt-get install protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Generate gRPC pkg

[`gen-protobuf.sh`](../grpc/gen-protobuf.sh)

List available `.proto`

```bash
$ protoc --decode_raw < dish.protoset | grep "\.proto\"" | grep "\"spacex" | awk '{print $2}' | sort | uniq

"spacex/api/common/status/status.proto"
"spacex/api/device/command.proto"
"spacex/api/device/common.proto"
"spacex/api/device/device.proto"
"spacex/api/device/dish_config.proto"
"spacex/api/device/dish.proto"
"spacex/api/device/transceiver.proto"
"spacex/api/device/wifi_config.proto"
"spacex/api/device/wifi.proto"
"spacex/api/device/wifi_util.proto"
"spacex/api/satellites/network/ut_disablement_codes.proto"
"spacex/api/telemetron/public/common/time.proto"
```

* Code Structure

* Add new metrics

* Add new panel and graph on Grafana Dashboard
