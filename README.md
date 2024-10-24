# DAPR PubSub - Go lang Example
An example of DAPR PubSub using Go lang. This is an implementation of producer consumer problem 

Other requirements
1. Configurable retry if the consumer fails to consume the message.
1. One consumer can concurrently process only one request

## Run pubsub program locally

## Prerequisites
1. Docker
1. Golang 1.19 or later
1. DAPR CLI 1.9 or later
1. Linux, Mac, or Windows (with WSL)

### Start rabbit mq
Using RabbitMQ as all the features are not supported in default redis.
```sh
docker run -d --restart unless-stopped --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4.0-management
```

### Initialize DAPR
```sh
dapr init 
```
### Run Consumers and Producers locally
Start Consumer-1
```sh
# Start DAPR Sidecar for Consumer-1
dapr run --app-id consumer1 --resources-path ./components --app-port 6002 --dapr-http-port 3602 --dapr-grpc-port 60002
# Start Consumer-1 application
go run consumer/main.go
```
Start Consumer-2
```sh
# Start DAPR Sidecar for Consumer-2
dapr run --app-id consumer2 --resources-path ./components --app-port 6003 --dapr-http-port 3603 --dapr-grpc-port 60003
# Start Consumer-2 application
go run consumer/main.go 6003
```
Start Producer-1
```sh
# Start DAPR Sidecar for Producer-1. Also starting he Producer-1 app in same command
dapr run --app-id producer --resources-path ./components --dapr-http-port 3601 --dapr-grpc-port 60001 go run producer/main.go
```

### Reference
1. [PubSub](https://docs.dapr.io/developing-applications/building-blocks/pubsub/howto-publish-subscribe/)
1. [Retry / Resiliency Policies](https://docs.dapr.io/operations/resiliency/policies/)
1. [RabbitMQ Configuration](https://docs.dapr.io/reference/components-reference/supported-pubsub/setup-rabbitmq/)
1. [Performance](https://docs.dapr.io/operations/performance-and-scalability/perf-service-invocation/)
