# CHEATSHEET

## Cluster and nodes
```bash
kubectl cluster-info
kubectl get nodes
```

## API and other resources
```bash
kubectl api-resources
kubectl get all --show-labels
```

## Config
```bash
kubectl config view
kubectl config get-contexts
```

## Namespaces
```bash
kubectl get namespaces --show-labels
kubectl create namespace my-namespace
kubectl apply -f ./namespaces/my-namespace.yml
# set default namespace
kubectl config set-context --current --namespace=my-namespace
```

## Pods
```bash
kubectl get pods --show-labels
# Show QoS of a pod
kubectl get pod my-pod --namespace=my-namespace --output jsonpath='{.status.qosClass}'
```

## Deployments
```bash
kubectl get deployment --show-labels
kubectl describe deployment my-nginx-deployment
kubectl apply -f ./deployments/nginx/.k8s
kubectl rollout status deployment my-nginx-deployment --watch
kubectl delete -f ./deployments/nginx/.k8s
```

## Jobs and Cron jobs
```bash
kubectl get jobs --show-labels
kubectl get cronjobs --show-labels
```

## Debugging
```bash
kubectl get deployments app --watch
kubectl logs deployments app
kubectl describe deployments app
kubectl exec --stdin --tty deployments/app -- /bin/sh
```

## Samples

### Redis pod and redis-cli
```bash
kubectl run my-redis --image=redis:6.2.6-alpine
kubectl port-forward pods/my-redis 6379:6379
# for redis-cli
kubectl exec my-redis -it  -- redis-cli
kubectl delete pod my-redis
```

### Excec to shell
````bash
kubectl exec --stdin --tty deployments/my-app -- /bin/sh
````
