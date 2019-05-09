# Kubernetes #74551 reproducer

## Requirements

* golang (preferrably 1.12+)
* running local Kubernetes cluster

## Run

Quickly build and deploy an http server in a local Kubernetes cluster:
```
./build_and_deploy.sh
```
Follow the server log in one terminal:
```
kubectl logs -f deploy/server
```
Then, in a different terminal run a client to poll the server via a port-forwarded connection:
```
./send_data.sh
```
