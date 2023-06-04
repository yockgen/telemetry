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
	Group  string
	Metric []string
	Host   []string
	Test   string
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
	var metrices, hosts string

	//metrices = "cpu_usage_idle|mem_used_percent"
	for i, itx := range qryObj.Metric {
		metrices = metrices + itx
		if i < len(qryObj.Metric)-1 {
			metrices = metrices + "|"
		}
	}

	//hosts = "TGL01|KBL01"
	for i, itx := range qryObj.Host {
		hosts = hosts + itx
		if i < len(qryObj.Host)-1 {
			hosts = hosts + "|"
		}
	}

	client := influxdb2.NewClient(influxURL, influxToken)
	defer client.Close()

	// Create a Flux query, TODO: need to modularized influx syntax
	fluxQuery = fmt.Sprintf(`from(bucket: "%s")
		|> range(start: -15m)
		|> filter(fn: (r) => r._measurement =~ /^%s$/ and r.host =~ /^%s$/)`, influxBucket, metrices, hosts)

	fmt.Println("QRY::", fluxQuery)

	// Execute the query
	queryAPI := client.QueryAPI(influxOrg)
	result, err := queryAPI.Query(context.Background(), fluxQuery)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return dataResult
	}
	defer result.Close()

	for result.Next() {

		record := result.Record()

		dataItx.Measurement = record.ValueByKey("_measurement").(string)
		dataItx.Value = record.ValueByKey("_value").(float64)
		dataItx.StartTsp = record.ValueByKey("_start").(time.Time)
		dataItx.StopTsp = record.ValueByKey("_stop").(time.Time)
		dataItx.Host = record.ValueByKey("host").(string)
		dataResult = append(dataResult, dataItx)

	}

	return dataResult

}

func rawtogrpc(dataResult []DataResult) *GetGroupResponse {

	var grpcResult = []*GetMetricResultResponse{}

	for _, itx := range dataResult {
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

	return result
}
