# Prerequisites

One liner install on Mac
```
> brew install kubectl helm minikube
```

### Detailed install guides if you're on a different platform
1. Install kubectl https://kubernetes.io/docs/tasks/tools/install-kubectl/
2. Install helm https://helm.sh/docs/intro/install/
3. Install minikube https://kubernetes.io/docs/tasks/tools/install-minikube/


# Quick Start

### Install

1. Start minikube
```
> ./scripts/minikube.sh
```
2. Check that you're kubectl is pointed at minikube
```
> kubectl config current-context
minikube
```
3. deploy application
```
> helm upgrade -i graphql-example chart/graphql-example --recreate-pods
```
4. Redeploy application
```
> helm del --purge graphql-example
> helm upgrade -i graphql-example chart/graphql-example --recreate-pods
```

### Pinging application

These steps allow you to send a request to the graphql application running inside minikube

1. In seperate tab run, this will "fake" a loadbalancer on your host machine and create a tunnel to the ingress of the application
```
> minikube tunnel
```
2. In main tab execute the follow curl curl
```
> curl -XPOST -d '{"query": "{ hello }"}' http://$(kubectl get svc -o jsonpath="{.items[0].spec.clusterIP}"):8080/graphql
{"data":{"hello":"Hello, world!"}}
```

