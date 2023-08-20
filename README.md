<p align="center">
  <h3 align="center">Starlink Prometheus Exporter Monitoring Stack</h3>
</p>

---

A [Starlink](https://www.starlink.com/) exporter for Prometheus. Not affiliated with or acting on behalf of Starlink(™)

[![goreleaser](https://github.com/clarkzjw/starlink_exporter/actions/workflows/release.yaml/badge.svg)](https://github.com/clarkzjw/starlink_exporter/actions/workflows/release.yaml)
[![build](https://github.com/clarkzjw/starlink_exporter/actions/workflows/build.yaml/badge.svg)](https://github.com/clarkzjw/starlink_exporter/actions/workflows/build.yaml)
[![License](https://img.shields.io/github/license/clarkzjw/starlink_exporter)](/LICENSE)
[![Release](https://img.shields.io/github/release/clarkzjw/starlink_exporter.svg)](https://github.com/clarkzjw/starlink_exporter/releases/latest)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/clarkzjw/starlink_exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/clarkzjw/starlink_exporter)](https://goreportcard.com/report/github.com/clarkzjw/starlink_exporter)

Original repository: 

+ https://github.com/danopstech/starlink
+ https://github.com/danopstech/starlink_exporter

---

gRPC last generated with firmware version: `87ed6ceb-4409-410a-b5f1-25886edf5966.uterm.release`

Firmware tracking website: https://starlinktrack.com/firmware/dishy

---

## Usage:

### Flags

`starlink_exporter` is configured through the use of optional command line flags

```bash
$ ./starlink_exporter --help
Usage of starlink_exporter
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

### Docker

Docker Images can be found at [Dockerhub](https://hub.docker.com/r/clarkzjw/starlink_exporter/tags).

Example:
```bash
docker pull docker.io/clarkzjw/starlink_exporter:latest

docker run \
  -p 9817:9817 \
  docker.io/clarkzjw/starlink_exporter:latest [flags]
```

### Setup Prometheus to scrape `starlink_exporter`

Configure [Prometheus](https://prometheus.io/) to scrape metrics from localhost:9817/metrics

```yaml
...
scrape_configs
    - job_name: starlink
      scrape_interval: 3s
      scrape_timeout:  3s
      static_configs:
        - targets: ['localhost:9817']
...
```

## Exported Metrics:

```text
# HELP starlink_dish_alert_install_pending Installation Pending
# TYPE starlink_dish_alert_install_pending gauge
starlink_dish_alert_install_pending 0
# HELP starlink_dish_alert_is_heating Is Heating
# TYPE starlink_dish_alert_is_heating gauge
starlink_dish_alert_is_heating 0
# HELP starlink_dish_alert_mast_not_near_vertical Status of mast position
# TYPE starlink_dish_alert_mast_not_near_vertical gauge
starlink_dish_alert_mast_not_near_vertical 0
# HELP starlink_dish_alert_motors_stuck Status of motor stuck
# TYPE starlink_dish_alert_motors_stuck gauge
starlink_dish_alert_motors_stuck 0
# HELP starlink_dish_alert_roaming Status of roaming
# TYPE starlink_dish_alert_roaming gauge
starlink_dish_alert_roaming 0
# HELP starlink_dish_alert_slow_eth_speeds Status of ethernet
# TYPE starlink_dish_alert_slow_eth_speeds gauge
starlink_dish_alert_slow_eth_speeds 0
# HELP starlink_dish_alert_thermal_shutdown Status of thermal shutdown
# TYPE starlink_dish_alert_thermal_shutdown gauge
starlink_dish_alert_thermal_shutdown 0
# HELP starlink_dish_alert_thermal_throttle Status of thermal throttling
# TYPE starlink_dish_alert_thermal_throttle gauge
starlink_dish_alert_thermal_throttle 0
# HELP starlink_dish_alert_unexpected_location Status of location
# TYPE starlink_dish_alert_unexpected_location gauge
starlink_dish_alert_unexpected_location 0
# HELP starlink_dish_anti_rollback_version Starlink Dish Anti Rollback Version.
# TYPE starlink_dish_anti_rollback_version counter
starlink_dish_anti_rollback_version 0
# HELP starlink_dish_boot_count Starlink Dish boot count.
# TYPE starlink_dish_boot_count counter
starlink_dish_boot_count 22
# HELP starlink_dish_bore_sight_azimuth_deg azimuth in degrees
# TYPE starlink_dish_bore_sight_azimuth_deg gauge
starlink_dish_bore_sight_azimuth_deg 3.60064435005188
# HELP starlink_dish_bore_sight_elevation_deg elevation in degrees
# TYPE starlink_dish_bore_sight_elevation_deg gauge
starlink_dish_bore_sight_elevation_deg 63.590576171875
# HELP starlink_dish_class_of_service User class of service
# TYPE starlink_dish_class_of_service gauge
starlink_dish_class_of_service{class_of_service="UNKNOWN_USER_CLASS_OF_SERVICE"} 1
# HELP starlink_dish_currently_obstructed Status of view of the sky
# TYPE starlink_dish_currently_obstructed gauge
starlink_dish_currently_obstructed 0
# HELP starlink_dish_dish_config Dish Config
# TYPE starlink_dish_dish_config gauge
starlink_dish_dish_config{level_dish_mode="TILT_LIKE_NORMAL",location_request_mode="LOCAL",power_save_mode="false",snow_melt_mode="AUTO"} 1
# HELP starlink_dish_dish_stow_requested stow requested
# TYPE starlink_dish_dish_stow_requested gauge
starlink_dish_dish_stow_requested 0
# HELP starlink_dish_downlink_throughput_bytes Amount of bandwidth in bytes per second download
# TYPE starlink_dish_downlink_throughput_bytes gauge
starlink_dish_downlink_throughput_bytes 42303.75
# HELP starlink_dish_eth_speed ethernet speed
# TYPE starlink_dish_eth_speed untyped
starlink_dish_eth_speed 1000
# HELP starlink_dish_first_nonempty_slot_seconds Seconds to next non empty slot
# TYPE starlink_dish_first_nonempty_slot_seconds gauge
starlink_dish_first_nonempty_slot_seconds 0
# HELP starlink_dish_fraction_obstruction_ratio Percentage of obstruction
# TYPE starlink_dish_fraction_obstruction_ratio gauge
starlink_dish_fraction_obstruction_ratio 0.09876814484596252
# HELP starlink_dish_gps_sats Number of GPS Sats.
# TYPE starlink_dish_gps_sats gauge
starlink_dish_gps_sats 14
# HELP starlink_dish_gps_valid GPS Status.
# TYPE starlink_dish_gps_valid gauge
starlink_dish_gps_valid 1
# HELP starlink_dish_info Running software versions and IDs of hardware
# TYPE starlink_dish_info gauge
starlink_dish_info{bootcount="22",country_code="CA",device_id="ut01000000-00000000-006d707e",hardware_version="rev3_proto2",manufactured_version="",software_version="87ed6ceb-4409-410a-b5f1-25886edf5966.uterm.release",utc_offset="-17999"} 1
# HELP starlink_dish_info_debug Debug Dish Info
# TYPE starlink_dish_info_debug gauge
starlink_dish_info_debug{count_by_reason="map[]",count_by_reason_delta="map[]",last_count="0",last_reason="BOOT_REASON_UNKNOWN"} 1
# HELP starlink_dish_is_dev Starlink Dish is Dev.
# TYPE starlink_dish_is_dev gauge
starlink_dish_is_dev 0
# HELP starlink_dish_is_hit Starlink Dish is Hit.
# TYPE starlink_dish_is_hit gauge
starlink_dish_is_hit 0
# HELP starlink_dish_location_info Dish Location Info (GPS/Starlink)
# TYPE starlink_dish_location_info gauge
starlink_dish_location_info{alt="70.97999999954364",lat="45.319438635558356",location_source="GPS",lon="-75.9437835"} 1
# HELP starlink_dish_mobility_class Dish mobility class
# TYPE starlink_dish_mobility_class gauge
starlink_dish_mobility_class{mobility_class="STATIONARY"} 1
# HELP starlink_dish_outage_did_switch Starlink Dish Outage Information
# TYPE starlink_dish_outage_did_switch gauge
starlink_dish_outage_did_switch 0
# HELP starlink_dish_outage_duration Starlink Dish Outage Information
# TYPE starlink_dish_outage_duration gauge
starlink_dish_outage_duration{cause="UNKNOWN",start_time="0"} 0
# HELP starlink_dish_pop_ping_drop_ratio Percent of pings dropped
# TYPE starlink_dish_pop_ping_drop_ratio gauge
starlink_dish_pop_ping_drop_ratio 0
# HELP starlink_dish_pop_ping_latency_seconds Latency of connection in seconds
# TYPE starlink_dish_pop_ping_latency_seconds gauge
starlink_dish_pop_ping_latency_seconds 0.03209523856639862
# HELP starlink_dish_prolonged_obstruction_duration_seconds Average in seconds of prolonged obstructions
# TYPE starlink_dish_prolonged_obstruction_duration_seconds gauge
starlink_dish_prolonged_obstruction_duration_seconds 2.9845340251922607
# HELP starlink_dish_prolonged_obstruction_interval_seconds Average prolonged obstruction interval in seconds
# TYPE starlink_dish_prolonged_obstruction_interval_seconds gauge
starlink_dish_prolonged_obstruction_interval_seconds 583.7838134765625
# HELP starlink_dish_prolonged_obstruction_valid Average prolonged obstruction is valid
# TYPE starlink_dish_prolonged_obstruction_valid gauge
starlink_dish_prolonged_obstruction_valid 0
# HELP starlink_dish_ready_state Dish ready states
# TYPE starlink_dish_ready_state gauge
starlink_dish_ready_state{aap="true",cady="true",l1l2="true",rf="true",scp="true",xphy="true"} 1
# HELP starlink_dish_scrape_duration_seconds Time to scrape metrics from starlink dish
# TYPE starlink_dish_scrape_duration_seconds gauge
starlink_dish_scrape_duration_seconds 0.161926935
# HELP starlink_dish_software_partitions_equal Starlink Dish Software Partitions Equal.
# TYPE starlink_dish_software_partitions_equal gauge
starlink_dish_software_partitions_equal 0
# HELP starlink_dish_up Was the last query of Starlink dish successful.
# TYPE starlink_dish_up gauge
starlink_dish_up 1
# HELP starlink_dish_uplink_throughput_bytes Amount of bandwidth in bytes per second upload
# TYPE starlink_dish_uplink_throughput_bytes gauge
starlink_dish_uplink_throughput_bytes 116165.4453125
# HELP starlink_dish_uptime_seconds Dish running time
# TYPE starlink_dish_uptime_seconds counter
starlink_dish_uptime_seconds 259142
# HELP starlink_dish_valid_seconds Unknown
# TYPE starlink_dish_valid_seconds counter
starlink_dish_valid_seconds 256210
```

## Example Grafana Dashboard:

[dashboard.json](./docker/config/grafana/provisioning/dashboards/Starlink.json)

<p align="center">
	<img src="https://github.com/clarkzjw/starlink_exporter/raw/main/.docs/assets/screenshot.png" width="95%">
</p>
