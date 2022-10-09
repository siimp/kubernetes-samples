## Setup
```bash
kubectl create namespace blue-green
kubectl config set-context --current --namespace=blue-green
```
## Service
```bash
kubectl apply -f service.yml
curl -s http://localhost:30000
```

## Blue
```bash
kubectl apply -f deployment-blue.yml
kubectl rollout status deployment my-nginx-blue --watch
```


## Green
```bash
kubectl apply -f deployment-green.yml
kubectl rollout status deployment my-nginx-green --watch
```


## Blue-Green Deployment
```bash
kubectl set selector service blue-green-service "app=my-nginx,role=green"
```
