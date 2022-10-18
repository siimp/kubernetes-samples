## Setup
```bash
kubectl delete namespace blue-green
kubectl create namespace blue-green
kubectl config set-context --current --namespace=blue-green
```
## Service
```bash
kubectl apply -f service.yml
go run requests.go
```

## Blue (v1.0.0)
```bash
kubectl apply -f deployment-blue.yml
kubectl rollout status deployment my-nginx-v1.0.0 --watch
```


## Green (v1.0.1)
```bash
kubectl apply -f deployment-green.yml
kubectl rollout status deployment my-nginx-v1.0.1 --watch
```


## Blue-Green Deployment
```bash
kubectl set selector service blue-green-service "app=my-nginx,role=v1.0.1"
kubectl delete deployment my-nginx-v1.0.0
```

## Info
For nginx configuration `keepalive_timeout 0` is neede for service switch to work. If keepalive timeout is greater than 0 switch wont work instantly
