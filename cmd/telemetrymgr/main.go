/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */
package main

import (
	"context"
	"log"

	"github.com/intel-innersource/infrastructure.edge.iaas.platform/internal/telemetrymgr"
	"github.com/spf13/cobra"
)

// This is our root command, to which we add subsequent commands, and subcommands
// in a tree-like fashion.
var rootCmd = &cobra.Command{
	Use:     "telemetrymgr",
	Short:   "telemetrymgr - Telemetry Manager for Intel Platform Director",
	Version: version,
	Long: `telemetrymgr is a service that provides access and configuraiton management
	to Intel Platform Director Telemetry Services.`,
	Run: func(cmd *cobra.Command, arg []string) {

	},
}

// Current telemetry-manager version - should be set during build
// e.g. -ldfalgs="-X" 'github.com/intel-innersource/.../main.go.version=x.y.z'" pdctl.go
const version = "0.0.1"
const configID = "telemetrymgr_config"

var (
	cfg string
)

func main() {

	var ctx = context.Background()
	var service telemetrymgr.TelemetrySvc

	rootCmd.Flags().StringVarP(&cfg, "config", "c", "telemetrymgr.yml", "Configuration file")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error, invalid usage %v", err)
	}

	ctx = context.WithValue(ctx, configID, cfg)
	service, err := telemetrymgr.New(ctx)
	if err != nil {
		log.Fatalf("Error, unable to create TelemetryMgr instance %v", err)
	}

	service.Start(ctx)

}
