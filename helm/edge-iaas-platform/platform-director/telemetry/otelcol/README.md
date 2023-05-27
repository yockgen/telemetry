### Prerequisites - setup environment variables 
```
source ./scripts/edge-iaas-platform/platform-director/telemetry/setenv.sh
```
Important Note: please review "setenv.sh" carefully, make change to ensure INTEL_TELEMETRY_PROJECT is pointing to the repo root directory. By default, the system assumed the repo will be clone under '/data/edgeiaas' parent directory, the full path of the repo is '/data/edgeiaas/{repo name}'

### 1. Deploy Otel Collector Backend Gateway 
```
helm install \
-f $INTEL_TELEMETRY_HELM/otelcol/values.yaml \
--set-file=otelconfig=$INTEL_TELEMETRY_HELM/otelcol/config/config-backend-gateway-01.yaml \
otelgateway $INTEL_TELEMETRY_HELM/otelcol
```

### 2. Verifying Backend Gateway 
```
# kubectl get svc
NAME                       TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)           AGE
kubernetes                 ClusterIP   10.233.0.1      <none>        443/TCP           2d16h
otelcol-gateway            ClusterIP   10.233.13.112   <none>        54199/TCP         32m
otelcol-gateway-nodeport   NodePort    10.233.4.81     <none>        54200:31082/TCP   32m
```

### 3. Pointing Otel Collector Frontend Agent to correct Backend Gateway 

```
nano $INTEL_TELEMETRY_HELM/otelcol/config/config-edge-agent.yaml
```
Modify following entry accordingly
```
exporters:
  otlp:    
    endpoint: http://{host ip address of Otel Collector Backend Gateway}:31082
```

### 4. Deploy Otel Collector Agent

```
source ./scripts/edge-iaas-platform/platform-manager/telemetry/setenv.sh
```

```
helm install \
-f $INTEL_TELEMETRY_HELM/otelcol/values-agent.yaml \
--set-file=otelconfig=$INTEL_TELEMETRY_HELM/otelcol/config/config-edge-agent.yaml \
otelcolagent $INTEL_TELEMETRY_HELM/otelcol
```

### 5. Verifying Frontend Agent 
```
# kubectl get svc
NAME                       TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)           AGE
kubernetes                 ClusterIP   10.233.0.1      <none>        443/TCP           2d16h
otelcol-agent-nodeport     NodePort    10.43.201.64    <none>        55200:30695/TCP   28s
otelcol-agent              ClusterIP   10.43.183.42    <none>        55199/TCP         28s
```


### 6. How to test the Open Telemetry infra from any telegraf node

Modified telegraf config to point to correct Otel Collector Agent
```
[[outputs.opentelemetry]]
  service_address  = "{host ip address of Otel Collector Frontend Agent}:30695"
```
