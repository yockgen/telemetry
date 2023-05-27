#!/bin/bash

export INTEL_TELEMETRY_ROOT="/data/edgeiaas"
export INTEL_TELEMETRY_PROJECT="$INTEL_TELEMETRY_ROOT/infrastructure.edge.iaas.platform-telemetry"
export INTEL_TELEMETRY_DOMAIN="platform-director"
export INTEL_TELEMETRY_HELM="$INTEL_TELEMETRY_PROJECT/helm/edge-iaas-platform/$INTEL_TELEMETRY_DOMAIN/telemetry"
export INTEL_TELEMETRY_SCRIPTS="$INTEL_TELEMETRY_PROJECT/scripts/edge-iaas-platform/$INTEL_TELEMETRY_DOMAIN/telemetry"
export INTEL_TELEMETRY_DOCKER="$INTEL_TELEMETRY_PROJECT/docker/edge-iaas-platform/$INTEL_TELEMETRY_DOMAIN/telemetry"

echo $INTEL_TELEMETRY_PROJECT
echo $INTEL_TELEMETRY_HELM
echo $INTEL_TELEMETRY_SCRIPTS
echo $INTEL_TELEMETRY_DOCKER


