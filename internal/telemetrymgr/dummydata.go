/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */
package telemetrymgr

// Static target metric result
var dummyTargetMetrices = []*GetMetricResponse{
	{
		Measurement: "mem_used_percent",
	},
	{
		Measurement: "cpu_usage_idle",
	},
	{
		Measurement: "powerstat_package_current_power_consumption_watts",
	},
}

// Static target metric result
var dummyTargetHosts = []*GetHostResponse{
	{
		Host: "TGL01",
	},
	{
		Host: "TGL02",
	},
	{
		Host: "KBL01",
	},
}

// Static sample metric result
var dummyMetrices = []*GetMetricResultResponse{
	{
		Measurement: "temperature",
		Host:        "server-1",
		Value:       25.5,
		Start:       "2023-05-30T10:00:00Z",
		Stop:        "2023-05-30T11:00:00Z",
	},
	{
		Measurement: "temperature",
		Host:        "server-2",
		Value:       25.5,
		Start:       "2023-05-30T10:00:00Z",
		Stop:        "2023-05-30T11:00:00Z",
	},
	{
		Measurement: "temperature",
		Host:        "server-3",
		Value:       25.5,
		Start:       "2023-05-30T10:00:00Z",
		Stop:        "2023-05-30T11:00:00Z",
	},
}

// Static sample data for an array of telemetry groups
var dummyGroups = []*GetGroupResponse{
	{
		Id:       "1",
		Name:     "cpu-core",
		Interval: 10,
		Latency:  5,
		Metric:   dummyTargetMetrices,
		Host:     dummyTargetHosts,
		//Result:   dummyMetrices,
	},
	{
		Id:       "2",
		Name:     "temp+cpu",
		Interval: 8,
		Latency:  2,
		Metric:   dummyTargetMetrices,
		Host:     dummyTargetHosts,
		//Result:   dummyMetrices,
	},
	{
		Id:       "3",
		Name:     "voltage+freq",
		Interval: 30,
		Latency:  60,
		Metric:   dummyTargetMetrices,
		Host:     dummyTargetHosts,
		//Result:   dummyMetrices,
	},
}
