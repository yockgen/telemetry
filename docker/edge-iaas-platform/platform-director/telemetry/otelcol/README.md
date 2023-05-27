### Prerequisites - setup environment variables 
```
source ./scripts/edge-iaas-platform/platform-director/telemetry/setenv.sh
```
Important Note: please review "setenv.sh" carefully, make change to ensure INTEL_TELEMETRY_PROJECT is pointing to the repo root directory. By default, the system assumed the repo will be clone under '/data/edgeiaas' parent directory, the full path of the repo is '/data/edgeiaas/{repo name}'

### Options of Deploying Open Telemetry Collector's with customized config file, for example:


Option 1: To Run Backend OtelCol Gateway with default config 
```
 CONFIG=$INTEL_TELEMETRY_DOCKER/otelcol/config-backend-gateway-01.yaml docker-compose up
```

Option 2: To Run Multiple OtelCol Gateways (refer next sections on setup load balancer to utlize multiple gateways)
```
 CONFIG=$INTEL_TELEMETRY_DOCKER/otelcol/config-backend-gateway-01.yaml docker-compose -f docker-compose-dual.yml up
```

Option 1: To Run Edge OtelCol Agent with default config 
```
source ./scripts/edge-iaas-platform/platform-manager/telemetry/setenv.sh
```
```
 CONFIG=$INTEL_TELEMETRY_DOCKER/otelcol/config-edge-agent.yaml docker-compose up
```


Please refer to example config files provided in this repo for reference.

### Setup Nginx Load Balancer 
For multiple OTELCOL gateways/agents setup on the cluster for better performance, NGINX load balancer could be use for this case, following demo setup is based on Ubuntu OS (please search respective NGINX LB setup online if user host OS is not Ubuntu):

1. Install Nginx
```
apt-get purge nginx nginx-common nginx-full
apt-get install nginx -y
```

2. Create load balancing configuration file 
```
nano /etc/nginx/conf.d/loadbalancer.conf
```

content looks like following:
```
upstream backend {        
        keepalive 16; 
        server {ip}:{port}; #<-- add in Otelcol instance 0, refer to respective docker compose file for port#
        server {ip}:{port}; #<-- add in Otelcol instance N, refer to respective docker compose file for port#    

    }

    server {
        listen 52123 http2;  #<-- user could defined any port not in used
        location / {	       
        	 grpc_pass grpc://backend;                                 
	}
 
}
```

3. Restart NGINX
```
systemctl restart nginx
systemctl status nginx.service
```

4. Configure Telegraf config file access to OtelCol cluster via Load Balancer, for example:

```
[[outputs.opentelemetry]]
  service_address  = "{LB host ip address}:52123"
```

