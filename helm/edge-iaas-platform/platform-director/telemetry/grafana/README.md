## Prerequisites 
1. Setup environment variables 
```
source ./scripts/edge-iaas-platform/platform-director/telemetry/setenv.sh
```
Important Note: please review "setenv.sh" carefully, make change to ensure INTEL_TELEMETRY_PROJECT is pointing to the repo root directory. By default, the system assumed the repo will be clone under '/data/edgeiaas' parent directory, the full path of the repo is '/data/edgeiaas/{repo name}'

2. Make Grafana Charts available in Helm 
```
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
```

## Configuring InfluxDB connection
1. Configuring InfluxDB connection in values.yaml
```
cd $INTEL_TELEMETRY_HELM/grafana
nano +525 values.yaml
```
2. Modify 'url' and 'token' field according to target InfluxDB (do not change other fields), for example:
```
## Configure grafana datasources
## ref: http://docs.grafana.org/administration/provisioning/#datasources
##
datasources:
  datasources.yaml:
    apiVersion: 1
    datasources:
      - name: Intel_Influx   
        uid: P8EFB3B3375746E67 
        type: influxdb
        access: proxy
        url: http://{ip address of influx's host}:32701
        jsonData:
          version: Flux
          organization: intel
          defaultBucket: intel
          tlsSkipVerify: true
        secureJsonData:
          token: {influxdb token for access}
```

* how to get port number from InfluxDB
```
# kubectl get svc | grep influx
intel-influxdb2            NodePort    10.233.52.128   <none>        80:32701/TCP      44h
```


## Grafana Helm deployment examples

### Option 1: Deployment without persistent storage - data wiped out after restart
```
helm install \
-f $INTEL_TELEMETRY_HELM/grafana/values.yaml \
--set adminPassword="admin" \
--set persistence.enabled=false \
--set persistence.storageClassName="manual" \
--set-file dashboards.default.power-insight-01.json=$INTEL_TELEMETRY_HELM/grafana/dashboards/power-insight.json \
--set-file dashboards.default.cluster-node-mon-01.json=$INTEL_TELEMETRY_HELM/grafana/dashboards/Cluster-Node-Monitoring.json \
intel-grafana grafana/grafana
```


### Option 2: Deployment with persistent storage

#### Pre - Cleanup - ensure environment is clean before deployment
```
cd $INTEL_TELEMETRY_HELM/grafana
helm uninstall intel-grafana
kubectl delete -f persistent.yaml
```
#### Pre - Re-create persistent storage 
```
kubectl apply -f persistent.yaml
```

#### Sample deployment 01
```
helm install \
-f $INTEL_TELEMETRY_HELM/grafana/values.yaml \
--set adminPassword="admin" \
--set persistence.enabled=true \
--set persistence.storageClassName="manual" \
intel-grafana grafana/grafana
```
#### Sample deployment 02 - import dashboard during deployment
```
helm install \
-f $INTEL_TELEMETRY_HELM/grafana/values.yaml \
--set adminPassword="admin" \
--set persistence.enabled=true \
--set persistence.storageClassName="manual" \
--set-file dashboards.default.power-insight-01.json=$INTEL_TELEMETRY_HELM/grafana/dashboards/power-insight.json \
--set-file dashboards.default.cluster-node-mon-01.json=$INTEL_TELEMETRY_HELM/grafana/dashboards/Cluster-Node-Monitoring.json \
intel-grafana grafana/grafana
```
### Reference
### 1. Grafana Helm Chart Official Guide
https://github.com/grafana/helm-charts/blob/main/charts/grafana/README.md

