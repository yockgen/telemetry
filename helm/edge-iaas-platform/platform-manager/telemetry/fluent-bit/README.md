### Prerequisites 

1. setup environment variables 
```
source ./scripts/edge-iaas-platform/platform-manager/telemetry/setenv.sh
```
Important Note: please review "setenv.sh" carefully, make change to ensure INTEL_TELEMETRY_PROJECT is pointing to the repo root directory. By default, the system assumed the repo will be clone under '/data/edgeiaas' parent directory, the full path of the repo is '/data/edgeiaas/{repo name}'

2.. Make Fluent Bit Charts available in Helm 
```
helm repo add fluent https://fluent.github.io/helm-charts
helm repo update
```

## Fluent Bit Helm deployment examples

### Deploy Fluent Bit with Example Common Configuration

1. Changing OUTPUT to desire Open Telemetry Collector (agent/node)
```
nano INTEL_TELEMETRY_HELM/fluent-bit/fluent-bit-outputs.yaml
```

Modifying host/IP and port accordingly
```
[OUTPUT]
    name opentelemetry
    match *
    host <ip/host of open telemetry collector, e.g. 10.158.76.160>
    port <port number of open telemetry collector, e.g. 31083>
    tls Off
    tls.verify Off

```

2. Deploy Fluent Bit using helm
```
helm install \
-f $INTEL_TELEMETRY_HELM/fluent-bit/values.yaml \
--set-file=config.inputs=$INTEL_TELEMETRY_HELM/fluent-bit/config/fluent-bit-inputs.conf \
--set-file=config.outputs=$INTEL_TELEMETRY_HELM/fluent-bit/config/fluent-bit-outputs.conf \
fluent-bit fluent/fluent-bit

```

3. Optional: adding other inputs    
Adding new inputs into ./config/fluent-bit-outputs.conf, re-run the helm chart
```
[INPUT]
    ....
    name tail    
    path /var/log/mylog.log
```

4. Optional: adding customized log's field, re-run the helm chart    
Modifying [FILTER] in ./values.yaml, for example, adding new field called 'host' in each row of log:
```
[FILTER]
    Name Lua
    Match *
    call append_tag
    code function append_tag(tag, timestamp, record) new_record = record new_record["host"] = "192.168.1.102".."@"..tag return 1, timestamp, new_record end    
```
More info: [https://docs.fluentbit.io/manual/pipeline/filters/lua](https://docs.fluentbit.io/manual/pipeline/filters/lua)

5. Clean up
```
helm uninstall fluent-bit
```
