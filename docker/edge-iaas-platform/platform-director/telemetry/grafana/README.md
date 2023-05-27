### Prerequisites 

1. setup environment variables 
```
source ./scripts/edge-iaas-platform/platform-director/telemetry/setenv.sh
```
Important Note: please review "setenv.sh" carefully, make change to ensure INTEL_TELEMETRY_PROJECT is pointing 
to the repo root directory. By default, the system assumed the repo will be clone under '/data/edgeiaas' 
parent directory, the full path of the repo is '/data/edgeiaas/{repo name}'

### Deploy grafana with default setting value

1. Deploy grafana on current terminal
```
cd $INTEL_TELEMETRY_DOCKER/grafana
docker-compose up
```

### Clean up grafana data with default setting
```
rm -r /var/lib/docker/volumes/grafana_grafana-storage/_data/*
```

### Deploy grafana with influxdb datasource and intel demo dashboard


create a datasource file
```
nano $INTEL_TELEMETRY_DOCKER/grafana/provisioning/datasource.yaml
```
Content of the file should look like this
```
apiVersion: 1

datasources:
  - name: Intel_Influx
    type: influxdb
    UID: P8EFB3B3375746E67 #<--- this has to match with datasource UID defined in $INTEL_TELEMETRY_DOCKER/provisioning/demo-dashboard.json
    access: proxy
    url: http://<Influxdb IP Address>:58100 #<--- insert influxdb IP
    jsonData:
      version: Flux
      organization: intel
      defaultBucket: intel
      tlsSkipVerify: true
    secureJsonData:
      token: _hnEeqaA9pZmHZpU4sroPmm-9VstwJekiQjuXbCVWk5ZuJ18gUCllTEkfKLZUNkOLYGkm0lMbvtFO-M4MkWBmA== #<--- influxdb access token
```

edit docker-compose.yml
```
cd $INTEL_TELEMETRY_DOCKER/grafana
nano docker-compose.yml
```
add provisioning folder mapping to volume
```
 volumes:
      - grafana-storage:/var/lib/grafana
      - ./provisioning/dashboards:/etc/grafana/provisioning/dashboards #<--- path mapping to your demo-dashboard.json and provider.yaml
      - ./provisioning/datasources:/etc/grafana/provisioning/datasources #<--- path mapping to your datasource.yaml
```
deploy grafana with datasource and dashboard loaded
```
cd $INTEL_TELEMETRY_DOCKER/grafana
docker-compose up
```

### Deploy grafana with user defined data storage path


edit docker-compose.yml
```
cd $INTEL_TELEMETRY_DOCKER/grafana
nano docker-compose.yml
```
add user id and change grafana-storage mapping to variable
```
    image: "grafana/grafana-enterprise:9.3.6"
    user: 1000:1000 #<--- add owner user ID
    volumes:
      - ${DATA}:/var/lib/grafana #<--- change to variable $DATA
```
create a directory for your data storage and set ownership to grafana user
```
mkdir /path/to/folder
chown -R 1000:1000 /path/to/folder
```
deploy grafana
```
cd $INTEL_TELEMETRY_DOCKER/grafana
DATA=/path/to/folder docker-compose up
```
### Clean up grafana data for custom setting
```
rm -r /path/to/folder/*
```

