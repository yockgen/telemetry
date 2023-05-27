### Prerequisites 

1. setup environment variables 
```
source ./scripts/edge-iaas-platform/platform-manager/telemetry/setenv.sh
```
Important Note: please review "setenv.sh" carefully, make change to ensure INTEL_TELEMETRY_PROJECT is pointing to the repo root directory. By default, the system assumed the repo will be clone under '/data/edgeiaas' parent directory, the full path of the repo is '/data/edgeiaas/{repo name}'

2.. Make Telegraf Charts available in Helm 
```
helm repo add influxdata https://helm.influxdata.com/
helm repo update
```

## Telegraf Helm deployment examples

### Deploy Telegraf with Example Common Configuration

1. Edit Configuration File
```
nano INTEL_TELEMETRY_HELM/telegraf/values.yaml
```
change open telemetry ip address and port
```
- opentelemetry:
    service_address: "<Node IP>:<Open Telemetry Port number>"
```

2. Deploy Telegraf using helm
```
cd INTEL_TELEMETRY_HELM/telegraf
kubectl apply -f persistent.yaml
helm install -f values.yaml intel-telegraf influxdata/telegraf
```

3. Clean up
```
cd INTEL_TELEMETRY_HELM/telegraf
helm uninstall intel-telegraf
kubectl delete -f persistent.yaml
```

### Deploy Telegraf with Example Intel Powerstat Configuration


1. Set user for Powerstat
```
chown -R 1000:1000 /sys/devices/virtual/powercap/intel-rapl
```

2. Edit Configuration File
```
nano INTEL_TELEMETRY_HELM/telegraf/powerstat/values.yaml
```
change open telemetry ip address and port
```
- opentelemetry:
    service_address: "<Node IP>:<Open Telemetry Port number>"
```

3. Deploy Telegraf using helm
```
cd INTEL_TELEMETRY_HELM/telegraf/powerstat/
kubectl apply -f persistent.yaml
helm install -f values.yaml intel-telegraf influxdata/telegraf
```

4. Clean up
```
cd INTEL_TELEMETRY_HELM/telegraf/powerstat/
helm uninstall intel-telegraf
kubectl delete -f persistent.yaml
```

### Deploy Telegraf with Example Intel RDT Configuration


1. Download and install RDT
```
git clone https://github.com/intel/intel-cmt-cat.git
cd intel-cmt-cat
make
sudo make install
```

2. Set user for RDT script
```
chown -R 1000:1000 /usr/local/bin/pqos
```

2. Edit Configuration File
```
nano INTEL_TELEMETRY_HELM/telegraf/rdt/values.yaml
```
change open telemetry ip address and port
```
- opentelemetry:
    service_address: "<Node IP>:<Open Telemetry Port number>"
```

3. Deploy Telegraf using helm
```
cd INTEL_TELEMETRY_HELM/telegraf/rdt/
kubectl apply -f persistent.yaml
helm install -f values.yaml intel-telegraf influxdata/telegraf
```

4. Clean up
```
cd INTEL_TELEMETRY_HELM/telegraf/rdt/
helm uninstall intel-telegraf
kubectl delete -f persistent.yaml
```


### Deploy Telegraf with Example Intel PMU Configuration


1. Download PMU tools to download event list
```
git clone https://github.com/andikleen/pmu-tools.git
cd pmu-tools
./event_download.py
```

2. Copy event list to user location
```
cp -r /root/.cache/pmu-events/ /home/user/pmu-events/
```

3. Set user for pmu event list
```
chown -R 1000:1000 /home/user/pmu-events/
```

4. Edit Configuration File
```
nano INTEL_TELEMETRY_HELM/telegraf/pmu/values.yaml
```
change open telemetry ip address and port
```
- opentelemetry:
    service_address: "<Node IP>:<Open Telemetry Port number>"
```

5. Deploy Telegraf using helm
```
cd INTEL_TELEMETRY_HELM/telegraf/pmu/
kubectl apply -f persistent.yaml
helm install -f values.yaml intel-telegraf influxdata/telegraf
```

6. Clean up
```
cd INTEL_TELEMETRY_HELM/telegraf/pmu/
helm uninstall intel-telegraf
kubectl delete -f persistent.yaml
```


### Deploy Telegraf with Example DPDK  Configuration


1. Install dpdk
```
sudo apt install dpdk
sudo apt-get install dpdk-dev libdpdk-dev
```

2. Run DPDK application, test-pmd
create hugepages
```
mkdir -p /dev/hugepages
mountpoint -q /dev/hugepages || mount -t hugetlbfs nodev /dev/hugepages
echo 2048 > /sys/devices/system/node/node0/hugepages/hugepages-2048kB/nr_hugepages
```
get the address of a NIC to use
```
dpdk-devbind.py -s
```
bind NIC to VFIO-PCI
```
dpdk-devbind.py -b=vfio-pci 58:00.0 #<--- change this address to your NIC address
```
run test-pmd
```
dpdk-testpmd --telemetry -- -i
```
leave the terminal open

3. On another terminal, set user to dpdk telemetry
```
chown -R 1000:1000 /var/run/dpdk/rte
```

4. Edit Configuration File
```
nano INTEL_TELEMETRY_HELM/telegraf/dpdk/values.yaml
```
change open telemetry ip address and port
```
- opentelemetry:
    service_address: "<Node IP>:<Open Telemetry Port number>"
```

5. Deploy Telegraf using helm
```
cd INTEL_TELEMETRY_HELM/telegraf/dpdk/
kubectl apply -f persistent.yaml
helm install -f values.yaml intel-telegraf influxdata/telegraf
```

6. Clean up
```
cd INTEL_TELEMETRY_HELM/telegraf/dpdk/
helm uninstall intel-telegraf
kubectl delete -f persistent.yaml
```


### Deploy Telegraf with Example RAS Configuration


1. Download RAS
```
apt install rasdaemon
```

2. Set user for RAS database
```
chown -R 1000:1000 /var/lib/rasdaemon
```

3. Edit Configuration File
```
nano INTEL_TELEMETRY_HELM/telegraf/ras/values.yaml
```
change open telemetry ip address and port
```
- opentelemetry:
    service_address: "<Node IP>:<Open Telemetry Port number>"
```

4. Deploy Telegraf using helm
```
cd INTEL_TELEMETRY_HELM/telegraf/ras/
kubectl apply -f persistent.yaml
helm install -f values.yaml intel-telegraf influxdata/telegraf
```

5. Clean up
```
cd INTEL_TELEMETRY_HELM/telegraf/ras/
helm uninstall intel-telegraf
kubectl delete -f persistent.yaml
```


### Deploy Telegraf with Example Kubernetes Configuration


1. Create user for telegraf, with admin rights
```
sudo kubectl create serviceaccount k8sadmin -n kube-system
sudo kubectl create clusterrolebinding k8sadmin --clusterrole=cluster-admin --serviceaccount=kube-system:k8sadmin
```

2. Create token for accessing kubernets API
```
kubectl create token k8sadmin -n kube-system --duration=999999h
```
save token to file
```
nano /home/user/k8telegraf/token
```

3. Set user for token file
```
chown -R 1000:1000 /home/user/k8telegraf/token
```

4. Edit Configuration File
```
nano INTEL_TELEMETRY_HELM/telegraf/kubernetes/values.yaml
```
change open telemetry ip address and port
```
- opentelemetry:
    service_address: "<Node IP>:<Open Telemetry Port number>"
```

5. Deploy Telegraf using helm
```
cd INTEL_TELEMETRY_HELM/telegraf/kubernetes/
kubectl apply -f persistent.yaml
helm install -f values.yaml intel-telegraf influxdata/telegraf
```

6. Clean up
```
cd INTEL_TELEMETRY_HELM/telegraf/kubernetes/
helm uninstall intel-telegraf
kubectl delete -f persistent.yaml
```

