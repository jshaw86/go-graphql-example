#!/bin/sh

minikube start --memory=4096 --cpus=4

helm init --wait

helm install --name prometheus stable/prometheus

helm install --name grafana stable/grafana

helm install --name mysql stable/mysql --set mysqlUser=some --set mysqlPassword=user --set mysqlDatabase=test

