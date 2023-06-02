/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */
package telemetrymgr

import (
	context "context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// influx connection
const (
	influxURL    = "http://192.168.1.107:32701"
	influxToken  = "X6zYQsXQdkC4K-WE7Uza_Z7yYWkENe3PAbNPIjryr4_KECA75QoLqALgsX9XQjWMFhdhZFz1TiLjxYUiM7B1zw=="
	influxOrg    = "intel"
	influxBucket = "intel"
)

type Query struct {
	UserId string `json:"id"`
	Token  string `json:"token"`
	Group  string `json:"group"`
	Test   string `json:"test"`
}

type DataResult struct {
	Measurement string    `json:"measurement"`
	Value       float64   `json:"value"`
	StartTsp    time.Time `json:"start"`
	StopTsp     time.Time `json:"stop"`
	Host        string    `json:"host"`
}

func retrieveData(qryObj Query) []DataResult {

	var dataItx DataResult
	var dataResult []DataResult
	var fluxQuery string

	metrices := "cpu_usage_idle"
	fmt.Println(metrices)

	client := influxdb2.NewClient(influxURL, influxToken)
	defer client.Close()

	// Create a Flux query, TODO: need to modularized influx syntax
	fluxQuery = fmt.Sprintf(`from(bucket: "%s")
		|> range(start: -15m)
		|> filter(fn: (r) => r._measurement == "%s")`, influxBucket, metrices)

	fmt.Println("QRY:: %s", fluxQuery)
	//return nil,nil

	// Execute the query
	queryAPI := client.QueryAPI(influxOrg)
	result, err := queryAPI.Query(context.Background(), fluxQuery)
	if err != nil {
		fmt.Printf("Error executing query:", err)
		return dataResult
	}
	defer result.Close()

	for result.Next() {

		//if result.TableChanged() {
		// New table started, print the table name
		//fmt.Printf("Table: %s\n", result.TableMetadata().String())
		//}

		record := result.Record()
		//val := record.ValueByKey("_value").(float64)

		dataItx.Measurement = record.ValueByKey("_measurement").(string)
		dataItx.Value = record.ValueByKey("_value").(float64)
		dataItx.StartTsp = record.ValueByKey("_start").(time.Time)
		dataItx.StopTsp = record.ValueByKey("_stop").(time.Time)
		dataItx.Host = record.ValueByKey("host").(string)
		dataResult = append(dataResult, dataItx)

	}

	return dataResult

}

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

// Retrieve the group list from persistent storage
func getGroups() (*GetGroupsResponse, error) {
	return &GetGroupsResponse{
		Groups: dummyGroups,
	}, nil
}

// Retrieve a specific group by index from storage
func getGroup(req *GetGroupRequest) (*GetGroupResponse, error) {

	var qry Query
	var rawResult []DataResult
	var grpcResult = []*GetMetricResultResponse{}

	if req == nil {
		return nil, nil
	}

	rawResult = retrieveData(qry)
	for _, itx := range rawResult {
		//fmt.Println("Measurement:", itx.Measurement)
		//fmt.Println("Value:", itx.Value)
		//fmt.Println("Start Time:", itx.StartTsp)
		//fmt.Println("Stop Time:", itx.StopTsp)
		//fmt.Println("Host:", itx.Host)
		//fmt.Println("----------------------")

		itxGrpc := &GetMetricResultResponse{
			Measurement: fmt.Sprintf("%s", itx.Measurement),
			Host:        fmt.Sprintf("%s", itx.Host),
			Value:       float64(itx.Value),
			Start:       fmt.Sprintf("%s", itx.StartTsp),
			Stop:        fmt.Sprintf("%s", itx.StopTsp),
		}
		grpcResult = append(grpcResult, itxGrpc)
	}

	var result = &GetGroupResponse{
		Id:       "1",
		Name:     "cpu-core",
		Interval: 10,
		Latency:  5,
		Result:   grpcResult,
	}

	size := len(grpcResult)
	fmt.Print("Size:", size)

	return result, nil

	//if req != nil {
	//if req.Index < uint32(len(dummyGroups)) {
	//	return dummyGroups[req.Index], nil
	//}
	//}
	//return nil, nil
}
