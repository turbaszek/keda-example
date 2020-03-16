# KEDA example

This repository consists of everything you need to setup simple Kubernetes 
cluster and demonstrate usage of KEDA redis and mysql scalers.

The included `helper` provides an easy way to perform both 0 -> n and n -> 0 scalings.  

## Create cluster
The deployment consists of 4 components:
- MySQL instance
- Redis instance
- Dummy pod that will be scaled up and down
- App service that provides some helper methods
```sh
kubectl -f apply deployment/
```

## Install KEDA
Follow the official KEDA guide https://keda.sh/deploy/


## Observe
To observe how everything works you can watch two things:
- number of pods and their state: `watch -n2 "kubectl get pods"`
- HPA stats: `watch -n2 "kubectl get hpa"`


## Redis example
To scale the dummy deployment using 
[Redis scaler](https://keda.sh/scalers/redis-lists/) first we have to
deploy the `ScaledObjects`:
```sh
kubectl -f apply keda/redis-hpa.yaml
```
this should result in creation of a new `ScaledObjects` and new HPA
```sh
# kubectl get scaledobjects
NAME                 DEPLOYMENT   TRIGGERS   AGE
redis-scaledobject   dummy        redis      5s

# kubectl get hpa
NAME             REFERENCE          TARGETS              MINPODS   MAXPODS   REPLICAS   AGE
keda-hpa-dummy   Deployment/dummy   <unknown>/10 (avg)   1         4         0          45s
```

To scale up we have to populate the Redis queue. To do this we can use the helper app:
```shell script
kubectl exec -it $(k get pods | grep "server" | cut -f 1 -d " ") app redis publish
```
and to scale down:
```shell script
kubectl exec -it $(k get pods | grep "server" | cut -f 1 -d " ") app redis drain
```

## MySQL example
To scale the dummy deployment using 
[MySQL scaler](https://keda.sh/scalers/mysql/) first we have to
deploy the `ScaledObjects`:
```sh
kubectl -f apply keda/mysql-hpa.yaml
```
this should result again in creation of `ScaleObject` and an HPA:
```sh
# kubectl get scaledobjects
NAME                 DEPLOYMENT   TRIGGERS   AGE
mysql-scaledobject   dummy        redis      5s

# kubectl get hpa
NAME             REFERENCE          TARGETS              MINPODS   MAXPODS   REPLICAS   AGE
keda-hpa-dummy   Deployment/dummy   <unknown>/10 (avg)   1         4         0          45s
```

To scale up we have to insert some values to MySQL database. 
To do this we can use the helper app:
```shell script
kubectl exec -it $(k get pods | grep "server" | cut -f 1 -d " ") app mysql insert
```
and to scale down:
```shell script
kubectl exec -it $(k get pods | grep "server" | cut -f 1 -d " ") app mysql delete
```
