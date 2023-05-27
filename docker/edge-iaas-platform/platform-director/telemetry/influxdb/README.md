### Prerequisites 

1. setup environment variables 
```
source ./scripts/edge-iaas-platform/platform-director/telemetry/setenv.sh
```
Important Note: please review 'setenv.sh' carefully, make change to ensure INTEL_TELEMETRY_PROJECT is pointing 
to the repo root directory. By default, the system assumed the repo will be clone under '/data/edgeiaas' 
parent directory, the full path of the repo is '/data/edgeiaas/{repo name}'

### Deploy influxdb with default setting value

1. Deploy influxdb on current terminal
```
cd $INTEL_TELEMETRY_DOCKER/influxdb
docker-compose up
```

### Clean up influxdb data with default setting
```
rm /var/lib/docker/volumes/influxdb_config/_data/influx-configs
rm -r /var/lib/docker/volumes/influxdb_data/_data/engine
rm /var/lib/docker/volumes/influxdb_data/_data/influxd.bolt
rm /var/lib/docker/volumes/influxdb_data/_data/influxd.sqlite
```

### Deploy influxdb with custom setting value

1. Deploy influxdb with user defined token


edit docker-compose.yml
```
cd $INTEL_TELEMETRY_DOCKER/influxdb
nano docker-compose.yml
```
define custom token by using DOCKER_INFLUXDB_INIT_ADMIN_TOKEN environment variable
```
 environment:
      - DOCKER_INFLUXDB_INIT_MODE=setup
      - DOCKER_INFLUXDB_INIT_USERNAME=admin
      - DOCKER_INFLUXDB_INIT_PASSWORD=intel@2023
      - DOCKER_INFLUXDB_INIT_ORG=intel
      - DOCKER_INFLUXDB_INIT_BUCKET=intel
      - DOCKER_INFLUXDB_INIT_RETENTION=8h
      # custom token
      - DOCKER_INFLUXDB_INIT_ADMIN_TOKEN=_hnEeqaA9pZmHZpU4sroPmm-9VstwJekiQjuXbCVWk5ZuJ18gUCllTEkfKLZUNkOLYGkm0lMbvtFO-M4MkWBmA==
    command: ["influxd"]

```
deploy influxdb with custom token
```
docker-compose up
```

2. Deploy influxdb with user defined path to store configuration and data


edit docker-compose.yml
```
cd $INTEL_TELEMETRY_DOCKER/influxdb
nano docker-compose.yml
```
define custom variable for config and data volume
```
volumes:
      - ${CONFIG}:/etc/influxdb2 #<--- config directory mapping
      - ${DATA}:/var/lib/influxdb2 #<--- data directory mapping
    ports:
```
deploy influxdb with custom volume
```
CONFIG=/path/to/config DATA=/path/to/data docker-compose up
```

### Clean up influxdb data for custom volume setting
```
rm -r /path/to/config
rm -r /path/to/data
```
