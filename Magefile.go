//go:build mage

/*
 * SPDX-FileCopyrightText: (C) 2023 Intel Corporation
 * SPDX-License-Identifier: LicenseRef-Intel
 */

package main

import (
	"context"
	"path/filepath"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	"os"
)

type Test mg.Namespace

// Runs Go tests.
func (Test) Go() error {
	// NOTE: Requires ginkgo v2 binary
	// TODO: Reintroduce -race detection once figuring out CGO with musl
	return sh.RunV("ginkgo", "--randomize-all", "--randomize-suites", "-v", "-r", "-tags", "unit", ".")
}

type Lint mg.Namespace

// Lints the REST schema.
func (Lint) Openapi() error {
	schemaPath, err := filepath.Abs("schemas/openapi.yaml")
	if err != nil {
		return err
	}

	schema, err := openapi3.NewLoader().LoadFromFile(schemaPath)
	if err != nil {
		return err
	}

	err = schema.Validate(context.Background())
	if err != nil {
		return err
	}

	return nil

}

type Telemetry mg.Namespace

var telemetryHelmPath = os.Getenv("INTEL_TELEMETRY_HELM")

// Open Telemetry Collector Helm validation
func (Telemetry) OtelValidate() error {

	name := "otelgateway"
	cmd := `helm template --dry-run --debug \
            -f ` + telemetryHelmPath + `/otelcol/values.yaml \
            --set-file=otelconfig=` + telemetryHelmPath + `/otelcol/config/config-backend-gateway-01.yaml \
            ` + name + ` ` + telemetryHelmPath + `/otelcol`

	return sh.RunV("bash", "-c", cmd)
}

// InfluxDB Helm Validation
func (Telemetry) InfluxValidate() error {

	name := "intel-influxdb2"
	cmd := `helm template --dry-run --debug \
            -f ` + telemetryHelmPath + `/influxdb2/values.yaml \
            --set persistence.enabled=false \
            --set adminUser.organization="intel" \
            --set adminUser.bucket="intel" \
            --set adminUser.token="X6zYQsXQdkC4K-WE7Uza_Z7yYWkENe3PAbNPIjryr4_KECA75QoLqALgsX9XQjWMFhdhZFz1TiLjxYUiM7B1zw==" \
            --set adminUser.user="admin" \
            --set adminUser.password="intel@2023" \
            ` + name + ` influxdata/influxdb2`

	return sh.RunV("bash", "-c", cmd)
}

// Grafana Helm Validation
func (Telemetry) GrafanaValidate() error {

	name := "intel-grafana"
	cmd := `helm template --dry-run --debug \
            -f ` + telemetryHelmPath + `/grafana/values.yaml \
            --set adminPassword="admin" \
            --set persistence.enabled=false \
            --set persistence.storageClassName="manual" \
            --set-file dashboards.default.power-insight-01.json=` + telemetryHelmPath + `/grafana/dashboards/power-insight.json \
            --set-file dashboards.default.cluster-node-mon-01.json=` + telemetryHelmPath + `/grafana/dashboards/Cluster-Node-Monitoring.json \
            ` + name + ` grafana/grafana`

	return sh.RunV("bash", "-c", cmd)
}
