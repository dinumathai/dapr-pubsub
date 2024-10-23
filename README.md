# DAPR PubSub - Go lang Example
An example of DAPR PubSub using Go lang.

## Prerequisites
1. Docker/nerdctl
1. Golang 1.19 or later
1. Dapr CLI 1.9 or later
1. Initialized Dapr environment
1. Linux, Mac, or Windows (with WSL)

## Run pubsub program locally

### Start rabbit mq
```
docker run -d --restart unless-stopped --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4.0-management
```

### Start jaeger for tracing
DO NOT USE `jaegertracing/all-in-one` FOR PRODUCTION
```
docker run -d --restart unless-stopped --name jaeger -e COLLECTOR_ZIPKIN_HOST_PORT=:9412 -p 16686:16686  -p 9412:9412 jaegertracing/all-in-one:1.22
```

### Run Consumer locally
```
dapr init 
```
```
dapr run --app-id consumer --resources-path ./components --app-port 6002 --dapr-http-port 3602 --dapr-grpc-port 60002 go run consumer/main.go
```
```
dapr run --app-id producer --resources-path ./components --app-port 6001 --dapr-http-port 3601 --dapr-grpc-port 60001 go run producer/main.go

```

### Reference
1. https://docs.dapr.io/developing-applications/building-blocks/pubsub/howto-publish-subscribe/
