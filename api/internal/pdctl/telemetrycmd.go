/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */
package pdctl

import (
	"github.com/spf13/cobra"
)

// This is the main telemetry command under which all telemetry subscommand
// and operations are supported
var telemetryCmd = &cobra.Command{
	Use:   "telemetry",
	Short: "Perform telemetry command",
	Long: `The telemetry subcommand enables full management of the telemetry-manager
functions in Platform Director which includes telemetry-groups, configuration 
and and policy management`,
}

func init() {
	rootCmd.AddCommand(telemetryCmd)
}
