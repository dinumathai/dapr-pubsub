# Run pubsub program in k8s

## Prerequisites
1. Kubernetes cluster with Rbac  up and running.
1. kubectl configured to access Kubernetes cluster with .
1. helm configured to access Kubernetes cluster.

## Docker build(optional)
The docker build is already done and pushed to docker-hub.
```sh
docker build -t dmathai/dapr-pubsub .
docker push dmathai/dapr-pubsub
```

## Install RabbitMQ
```sh
# TODO - durable/persistent storage
helm install rabbitmq ./rabbitmq-15.0.5.tgz --namespace rabbitmq --create-namespace --wait -f rmq-values.yaml
```

## Install dapr
```sh
helm install dapr ./dapr-1.14.4.tgz --namespace dapr-system --create-namespace --wait -f dapr-values.yaml
```

## Deploy Consumer and Producer app
```sh
# Creating producer and consumer on default namespace
kubectl apply -k resources/
```

## Uninstall all components
```sh
kubectl delete -k resources/

helm uninstall dapr -n dapr-system
kubectl delete ns dapr-system

helm uninstall rabbitmq -n rabbitmq
kubectl delete ns rabbitmq
```

## Reference
1. [Helm Charts](https://github.com/dapr/helm-charts)
1. [Helm Chart documentation](https://github.com/dapr/dapr/blob/master/charts/dapr/README.md)

