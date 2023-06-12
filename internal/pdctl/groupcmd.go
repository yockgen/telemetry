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

// This is the main collections command under which all collection sub-commands
// and operations are supported.
var collectionCmd = &cobra.Command{
	Use:   "collections",
	Short: "Executes commands based on a telemetry collection",
	Long: `Manages telemetry collections enabling creation, updating, deleting and
retreving details about a telemetry collection. Telemetry collections are
managed on a per tenant basis in platform director isolated across users.`,
	Aliases: []string{"collection"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// This is the command to enumerate, list all the collections.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the telemetry collections",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := GetTelemetryCollections()
		if err == nil {
			fmt.Printf("%+v", list)
		}
	},
}

// This is the command to retrieve a specifc collection by ID.
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an instance of a telemetry collection",
	Run: func(cmd *cobra.Command, args []string) {

		idx, _ := cmd.Flags().GetUint32("id")
		collection, err := GetTelemetryCollection(idx)
		if err == nil {
			fmt.Printf("test %+v", collection)
		}
	},
}

// Retrieves all telemetry collections from the Telemetry Manager
func GetTelemetryCollections() (*tm.GetCollectionsResponse, error) {

	list, err := getTelemetryGrpcClient().GetCollections(context.Background(), &tm.GetCollectionsRequest{})
	if nil != err {
		fmt.Printf("Error getting the list of collections %v", err)
	}
	return list, err
}

// Retrieves a telemetry collection instance by ID from Telemetry Manager
func GetTelemetryCollection(idx uint32) (*tm.GetCollectionResponse, error) {

	collection, err := getTelemetryGrpcClient().GetCollection(context.Background(), &tm.GetCollectionRequest{Index: idx})

	if nil != err {
		fmt.Printf("Error getting collection instance %v", err)
	}

	return collection, err
}

// Init function is used to to build the command hierarchy and configuring
// flags - global and local to context
func init() {

	telemetryCmd.AddCommand(collectionCmd)
	collectionCmd.AddCommand(listCmd)
	collectionCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().Uint32("id", 0, "Identifies a telemetry collection instance")
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
