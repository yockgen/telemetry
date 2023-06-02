/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */
package pdctl

import (
	"context"
	"fmt"
	"log"

	tm "github.com/intel-innersource/infrastructure.edge.iaas.platform/internal/telemetrymgr"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// This is the main group command under which all groups sub-commands
// and operations are supported.
var groupCmd = &cobra.Command{
	Use:   "groups",
	Short: "Executes commands based on a telemetry group",
	Long: `Manages telemetry groups enabling creation, updating, deleting and
retreving details about a telemetry group. Telemetry groups are
managed on a per tenant basis in platform director isolated across users.`,
	Aliases: []string{"group"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// This is the command to enumerate, list all the groups.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the telemetry groups",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := GetTelemetryGroups()
		if err == nil {
			fmt.Printf("%+v", list)
		}
	},
}

// This is the command to retrieve a specifc group by ID.
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an instance of a telemetry group",
	Run: func(cmd *cobra.Command, args []string) {

		idx, _ := cmd.Flags().GetUint32("id")
		group, err := GetTelemetryGroup(idx)
		if err == nil {
			fmt.Printf("test %+v", group)
		}
	},
}

// Retrieves all telemetry groups from the Telemetry Manager
func GetTelemetryGroups() (*tm.GetGroupsResponse, error) {

	list, err := getTelemetryGrpcClient().GetGroups(context.Background(), &tm.GetGroupsRequest{})
	if nil != err {
		fmt.Printf("Error getting the list of groups %v", err)
	}
	return list, err
}

// Retrieves a telemetry group instance by ID from Telemetry Manager
func GetTelemetryGroup(idx uint32) (*tm.GetGroupResponse, error) {

	group, err := getTelemetryGrpcClient().GetGroup(context.Background(), &tm.GetGroupRequest{Index: idx})

	if nil != err {
		fmt.Printf("Error getting group instance %v", err)
	}

	return group, err
}

// Init function is used to to build the command hierarchy and configuring
// flags - global and local to context
func init() {

	telemetryCmd.AddCommand(groupCmd)
	groupCmd.AddCommand(listCmd)
	groupCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().Uint32("id", 0, "Identifies a telemetry group instance")
}

func getTelemetryGrpcClient() tm.TelemetryServiceExternalClient {

	// TODO: configuration from a local yml
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to create grpc instance %v", err)
		return nil
	}
	return tm.NewTelemetryServiceExternalClient(conn)
}
