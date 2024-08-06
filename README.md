# kubernetes-samples

## How to run kubernetes locally

### Docker and Docker Desktop
Install docker desktop and enable kubernetes via gui.

### Podman and minikube

1. run podman in rootless mode
```
podman machine set --rootful=false
podman machine start
```
2. install minikube (rootless podman)
```
# https://minikube.sigs.k8s.io/docs/start/
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
minikube config set rootless true
minikube config set driver podman
minikube config set container-runtime containerd
minikube start --kubernetes-version=v1.29.7 --memory 4096 --cpus=4
```
```
# if cpu delegate fails
# https://rootlesscontaine.rs/getting-started/common/cgroup2/
sudo mkdir -p /etc/systemd/system/user@.service.d
cat <<EOF | sudo tee /etc/systemd/system/user@.service.d/delegate.conf
[Service]
Delegate=cpu cpuset io memory pids
EOF
sudo systemctl daemon-reload
```
```
# Manage cluster via k9s
podman run --rm -it --network=host -v ~/.kube/config:/root/.kube/config -v ~/.minikube:$(echo ~)/.minikube docker.io/derailed/k9s:v0.27.4
```





## Other resources
https://kubernetes.io/docs/reference/kubectl/cheatsheet/  
https://github.com/kubernetes/examples  
