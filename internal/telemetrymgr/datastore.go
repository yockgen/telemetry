/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */
package telemetrymgr

// Retrieve the collection list from persistent storage
func getCollections() (*GetCollectionsResponse, error) {
	return &GetCollectionsResponse{
		Collections: dummyCollections,
	}, nil
}

// Retrieve a collection from persistent storage, this is a dummy function, TODO: get from database
func _getCollection(idx uint32) *GetCollectionResponse {

	if idx < uint32(len(dummyCollections)) {
		return dummyCollections[idx]
	}
	return nil
}

// Retrieve a specific collection by index from storage
func getCollection(req *GetCollectionRequest) (*GetCollectionResponse, error) {

	var qry Query
	var rawResult []DataResult
	var result *GetCollectionResponse

	if req == nil {
		return nil, nil
	}

	//var reqParam *GetCollectionResponse = dummyCollections[req.Index]
	reqParam := _getCollection(req.Index)

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
