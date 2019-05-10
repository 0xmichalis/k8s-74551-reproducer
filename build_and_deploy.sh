#!/usr/bin/env bash

set -ex

GOOS=linux go build -o server server.go

docker build -t test-server:test .

kubectl create ns repro74551 || true
kubectl delete deploy/server -n repro74551 >/dev/null 2>&1 || true
kubectl run server --image test-server:test -n repro74551
kubectl patch deploy/server -n repro74551 -p '{"spec":{"template":{"spec":{"containers":[{"name":"server","ports":[{"containerPort":8080}]}]}}}}'
kubectl rollout status deploy/server -n repro74551

