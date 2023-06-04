/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */
package telemetrymgr

// Retrieve the group list from persistent storage
func getGroups() (*GetGroupsResponse, error) {
	return &GetGroupsResponse{
		Groups: dummyGroups,
	}, nil
}

// Retrieve a group from persistent storage, this is a dummy function, TODO: get from database
func _getGroup(idx uint32) *GetGroupResponse {

	if idx < uint32(len(dummyGroups)) {
		return dummyGroups[idx]
	}
	return nil
}

// Retrieve a specific group by index from storage
func getGroup(req *GetGroupRequest) (*GetGroupResponse, error) {

	var qry Query
	var rawResult []DataResult
	var result *GetGroupResponse

	if req == nil {
		return nil, nil
	}

	//var reqParam *GetGroupResponse = dummyGroups[req.Index]
	reqParam := _getGroup(req.Index)

	//grpc message conversion to customized influx param structure
	for _, param := range reqParam.Metric {
		//fmt.Println("metrices=", param.Measurement)
		qry.Metric = append(qry.Metric, param.Measurement)

	}

	for _, param := range reqParam.Host {
		//fmt.Println("Host=", param.Host)
		qry.Host = append(qry.Host, param.Host)

	}

	//query influxdb
	rawResult = retrieveData(qry)
	result = rawtogrpc(rawResult)

	return result, nil

}
