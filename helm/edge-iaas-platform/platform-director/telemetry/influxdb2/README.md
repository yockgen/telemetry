### Prerequisites 

1. setup environment variables 
```
source ./scripts/edge-iaas-platform/platform-director/telemetry/setenv.sh
```
Important Note: please review "setenv.sh" carefully, make change to ensure INTEL_TELEMETRY_PROJECT is pointing to the repo root directory. By default, the system assumed the repo will be clone under '/data/edgeiaas' parent directory, the full path of the repo is '/data/edgeiaas/{repo name}'

2.. Make InfluxDB2 Charts available in Helm 
```
helm repo add influxdata https://helm.influxdata.com/
helm repo update
```

## InfluxDB Helm deployment examples

### Option 1: Deployment with persistent storage

1. Cleanup previous deployment and persistent volume
```
helm uninstall intel-influxdb2
cd $INTEL_TELEMETRY_HELM/influxdb2
kubectl delete -f persistent.yml
```
2. Deploy persistent volume
```
cd $INTEL_TELEMETRY_HELM/
kubectl apply -f persistent.yml
```
3. Deploy with persistent storage - data retained
```
helm install \
-f $INTEL_TELEMETRY_HELM/influxdb2/values.yaml \
--set persistence.enabled=true,persistence.useExisting=true \
--set persistence.existingClaim=intel-influxdb2 \
--set persistence.storageClass="-" \
--set adminUser.organization="intel" \
--set adminUser.bucket="intel" \
--set adminUser.token="X6zYQsXQdkC4K-WE7Uza_Z7yYWkENe3PAbNPIjryr4_KECA75QoLqALgsX9XQjWMFhdhZFz1TiLjxYUiM7B1zw==" \
--set adminUser.user="admin" \
--set adminUser.password="intel@2023" \
intel-influxdb2 influxdata/influxdb2
```

### Option 2: Deployment without persistent storage - data wiped out after restart

```
helm install \
-f $INTEL_TELEMETRY_HELM/influxdb2/values.yaml \
--set persistence.enabled=false \
--set adminUser.organization="intel" \
--set adminUser.bucket="intel" \
--set adminUser.token="X6zYQsXQdkC4K-WE7Uza_Z7yYWkENe3PAbNPIjryr4_KECA75QoLqALgsX9XQjWMFhdhZFz1TiLjxYUiM7B1zw==" \
--set adminUser.user="admin" \
--set adminUser.password="intel@2023" \
intel-influxdb2 influxdata/influxdb2
```

### Reference

### 1. InfluxDB Helm Chart Official Guide
[https://github.com/grafana/helm-charts/blob/main/charts/grafana/README.md](https://github.com/influxdata/helm-charts/blob/master/charts/influxdb2/README.md)

