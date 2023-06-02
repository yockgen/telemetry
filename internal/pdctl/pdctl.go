/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */
package pdctl

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Current CLI version - should be set during build
// e.g. -ldfalgs="-X" 'github.com/intel-innersource/.../pdctl.version=x.y.z'" pdctl.go
const version = "0.0.1"

// Global verbose logging flags
var verbose bool

// This is our root command, to which we add subsequent commands, and subcommands
// in a tree-like fashion.
var rootCmd = &cobra.Command{
	Use:     "pdctl",
	Short:   "pdctl - a CLI for Intel Platform Director",
	Version: version,
	Args:    cobra.MinimumNArgs(1),
	Long: `pdctl is a CLI to access and manage different instances of Platform Director
for on-premise and cloud hosted instances. The client can execute commands
and subcommands defined in the list below`,
	Run: func(cmd *cobra.Command, arg []string) {

	},
}

func Execute() {

	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error, invalid usage")
	}
}
