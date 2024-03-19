- Get pods of Pending and Running status

```
kubectl get pods -n dev --field-selector=status.phase=Pending
kubectl get pods -n dev --field-selector=status.phase=Running
```
- Scale RepicationController

```
kubectl scale --replicas=1000 replicationcontroller/nginx-rc -n dev
```

- Describe a replication controller

```
kubectl describe -n dev replicationcontroller/nginx-rc
```

- Get Node Information

```
kubectl get nodes -o wide
```

- Add new labels to the node

```
kubectl label nodes minikube-m03 workload=dev colour=yellow
```

- kubectl port-forward service

```
kubectl port-forward service/uploadapp-service 8082:8081 -n dev
```
- kubectl port-forward pod

```
kubectl port-forward pod/uploadapp-service-abcd-234 8082:8081 -n dev
```

- create a configmap using command

```
kubectl create configmap demo-confgmap5 --from-file=configmap-file.json -n dev
```

- create a configmap using keys and values

```
kubectl create configmap demo-configmap6 -n dev --from-literal=NAME=Hitachi --from-literal=TRAINING=Devops
```