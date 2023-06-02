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

// Retrieve a specific group by index from storage
func getGroup(req *GetGroupRequest) (*GetGroupResponse, error) {

	var qry Query
	var rawResult []DataResult
	var result *GetGroupResponse

	if req == nil {
		return nil, nil
	}

	rawResult = retrieveData(qry)
	result = rawtogrpc(rawResult)

	return result, nil

	//if req != nil {
	//	if req.Index < uint32(len(dummyGroups)) {
	///return dummyGroups[req.Index], nil
	//}
	//}
	//return nil, nil

}
