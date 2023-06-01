/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */
package telemetrymgr

// Static target metric result
var dummyTargetMetrices = []*GetMetricResponse{
	{
		Measurement: "temperature",
	},
	{
		Measurement: "cpu_idle",
	},
	{
		Measurement: "temperature",
	},
}

// Static target metric result
var dummyTargetHosts = []*GetHostResponse{
	{
		Host: "server-1",
	},
	{
		Host: "server-2",
	},
	{
		Host: "server-3",
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
		Result:   dummyMetrices,
	},
	{
		Id:       "2",
		Name:     "temp+cpu",
		Interval: 8,
		Latency:  2,
		Metric:   dummyTargetMetrices,
		Host:     dummyTargetHosts,
		Result:   dummyMetrices,
	},
	{
		Id:       "3",
		Name:     "voltage+freq",
		Interval: 30,
		Latency:  60,
		Metric:   dummyTargetMetrices,
		Host:     dummyTargetHosts,
		Result:   dummyMetrices,
	},
}

// Retrieve the group list from persistent storage
func getGroups() (*GetGroupsResponse, error) {
	return &GetGroupsResponse{
		Groups: dummyGroups,
	}, nil
}

// Retrieve a specific group by index from storage
func getGroup(req *GetGroupRequest) (*GetGroupResponse, error) {

	if req != nil {
		if req.Index < uint32(len(dummyGroups)) {
			return dummyGroups[req.Index], nil
		}
	}
	return nil, nil
}
