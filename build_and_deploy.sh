#!/usr/bin/env bash

set -ex

GOOS=linux go build -o server server.go

docker build -t test-server:test .

kubectl delete deploy/server >/dev/null 2>&1 || true
kubectl run server --image test-server:test
kubectl patch deploy/server -p '{"spec":{"template":{"spec":{"containers":[{"name":"server","ports":[{"containerPort":8080}]}]}}}}'
kubectl rollout status deploy/server

