#!/usr/bin/env bash

kubectl port-forward deploy/server 8080:8080 &

until nc -z localhost 8080 >/dev/null 2>&1; do
   sleep 0.1
done

go run client.go
