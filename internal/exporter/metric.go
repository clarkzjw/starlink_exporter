package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Location Info
	dishLocationInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "location_info"),
		"Dish Location Info (GPS/Starlink)",
		[]string{
			"location_source",
			"lat",
			"lon",
			"alt"}, nil,
	)

	// DeviceInfo
	dishInfo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "info"),
		"Running software versions and IDs of hardware",
		[]string{
			"device_id",
			"build_id",
			"hardware_version",
			"software_version",
			"generationNumber",
			"country_code",
			"bootcount",
			"utc_offset"}, nil,
	)
	dishInitializationDurationSeconds = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "initialization_duration_seconds"),
		"Initialization duration in seconds",
		[]string{
			"attitudeInitialization",
			"burstDetected",
			"ekfConverged",
			"firstCplane",
			"firstPopPing",
			"gpsValid",
			"initialNetworkEntry",
			"networkSchedule",
			"rfReady",
			"stableConnection",
		}, nil,
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
	dishAlignmentStats = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alignment_stats"),
		"Starlink Dish Alignment Stats",
		[]string{
			"tiltAngleDeg",
			"boresightAzimuthDeg",
			"boresightElevationDeg",
			"attitudeEstimationState",
			"attitudeUncertaintyDeg",
			"desiredBoresightAzimuthDeg",
			"desiredBoresightElevationDeg"}, nil,
	)
	// DeviceState
	dishUp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "up"),
		"Was the last query of Starlink dish successful.",
		nil, nil,
	)
	// dishScrapeDurationSeconds = prometheus.NewDesc(
	// 	prometheus.BuildFQName(namespace, "dish", "scrape_duration_seconds"),
	// 	"Time to scrape metrics from starlink dish",
	// 	nil, nil,
	// )
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
	dishSoftwareUpdateState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "software_update_state"),
		"Starlink Dish Software Update State",
		[]string{"software_update_state"}, nil,
	)
	dishDisablementCode = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "disablement_code"),
		"Starlink Dish Disablement Code",
		[]string{"disablement_code"}, nil,
	)
	// DishGpsStats
	dishGpsValid = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "gps_valid"),
		"GPS Status",
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
	dishPhyRxBeamSnrAvg = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "phy_rx_beam_snr_avg"),
		"physical rx beam snr average",
		nil, nil,
	)
	dishTemperateCenter = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "tCenter"),
		"Temperature center",
		nil, nil,
	)
	dishEthSpeedMbps = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "eth_speed"),
		"ethernet speed",
		nil, nil,
	)
	// DishAlerts
	dishPowerSupplyThermalThrottle = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_power_supply_thermal_throttle"),
		"Status of power supply thermal throttling",
		nil, nil,
	)
	dishIsPowerSaveIdle = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_power_save_idle"),
		"Status of power save idle",
		nil, nil,
	)
	dishMovingWhileNotMobile = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_moving_while_not_mobile"),
		"Status of moving while not mobile",
		nil, nil,
	)
	dishMovingTooFastForPolicy = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_moving_too_fast_for_policy"),
		"Status of moving too fast for policy",
		nil, nil,
	)
	dishLowMotorCurrent = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_low_motor_current"),
		"Status of low motor current",
		nil, nil,
	)
	dishLowerSignalThanPredicted = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_lower_signal_than_predicted"),
		"Status of lower signal than predicted",
		nil, nil,
	)
	dishObstructionMapReset = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "alert_obstruction_map_reset"),
		"Status of obstruction map reset",
		nil, nil,
	)
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
	dishPatchesValid = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "patches_valid"),
		"Number of valid patches",
		nil, nil,
	)
	dishCurrentlyObstructed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "currently_obstructed"),
		"Status of view of the sky",
		nil, nil,
	)
	dishTimeObstructed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "time_obstructed"),
		"Time obstructed ratio",
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

	// dishObstructionMap
	dishObstructionMap = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "obstruction_map"),
		"Obstruction Map",
		[]string{
			"timestamp",
			"num_rows",
			"num_cols",
			// "min_elevation_deg",
			"max_theta_deg",
			"reference_frame",
			"data"}, nil,
	)

	// diagnostics
	dishGpsTimeS = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "gps_time_s"),
		"GPS Time",
		nil, nil,
	)
	// TODO:
	// Find a Golang package to convert GPS time to UTC time
	// dishUTCTime = prometheus.NewDesc(
	// 	prometheus.BuildFQName(namespace, "dish", "utc_time"),
	// 	"UTC Time",
	// 	nil, nil,
	// )

	// dishPowerWatt
	dishPowerWatt = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "power_watt_current"),
		"Current Power Usage in Watt",
		nil, nil,
	)
	dishPowerWattAvg15min = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "dish", "power_watt_avg_15min"),
		"Average Power Usage in Watt over 15 minutes",
		nil, nil,
	)
)
