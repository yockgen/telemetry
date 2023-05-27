
### Prerequisites

1. setup environment variables
```
source ./scripts/edge-iaas-platform/platform-director/telemetry/setenv.sh
```
Important Note: please review 'setenv.sh' carefully, make change to ensure INTEL_TELEMETRY_PROJECT is pointing
to the repo root directory. By default, the system assumed the repo will be clone under '/data/edgeiaas'
parent directory, the full path of the repo is '/data/edgeiaas/{repo name}'

### Deploy Fluent-bit

1. Deploy Fluent-bit on current terminal
```
cd $INTEL_TELEMETRY_DOCKER/fluentbit
CONFIG=./path/to/config docker-compose up
```

### Deploy Fluent-bit with Example Common Configuration


1. Edit Configuration File
```
nano INTEL_TELEMETRY_DOCKER/fluentbit/configuration/fluent-bit-common.conf
```
change open telemetry ip address and port
```
[OUTPUT]
    name opentelemetry
    match *
    host <Node IP>
    port <Open Telemetry HTTP Port number>
```

2. Deploy Fluent-bit on current terminal
```
cd $INTEL_TELEMETRY_DOCKER/fluentbit
CONFIG=./configuration/fluent-bit-common.conf docker-compose up
```
