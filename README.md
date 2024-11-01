# DAPR PubSub - Go lang Example
An example of DAPR PubSub using Go lang. This is an implementation of producer consumer problem 

__Requirements of Example 1 - Using DAPR golang SDK__
1. Configurable retry if the consumer fails to consume the message.
1. One consumer can concurrently process only one request

__Requirements of Example 2 - Using plain http apis__
1. Example 1 Requirements.
1. Support Raw messages without DAPR CloudEvents wrapper.

## Run pubsub program in k8s
Refer [doc](./k8s/README.md)

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
__Start Consumer-1__
```sh
# Start DAPR Sidecar for Consumer-1
dapr run --app-id consumer --resources-path ./components --app-port 6002 --dapr-http-port 3602
```
```sh
# Start Consumer-1 application that uses dapr sdk
go run dapr-lib/consumer/main.go 6002
```
OR 
```sh
# Start Consumer-1 application that uses plain http
go run http/consumer/main.go 6002
```
__Start Consumer-2__
```sh
# Start DAPR Sidecar for Consumer-2
dapr run --app-id consumer --resources-path ./components --app-port 6003 --dapr-http-port 3603
```
```sh
# Start Consumer-2 application that uses dapr sdk
go run dapr-lib/consumer/main.go 6003
```
OR 
```sh
# # Start Consumer-2 application that uses plain http
go run http/consumer/main.go 6003
```
__Start Producer__
```sh
# Start DAPR Sidecar for Producer-1. Also starting he Producer-1 app(that uses dapr sdk) in same command
dapr run --app-id producer --resources-path ./components --dapr-grpc-port 60001 go run dapr-lib/producer/main.go 60001
```

OR

```sh
# Start DAPR Sidecar for Producer-1. Also starting he Producer-1 app(that uses plain http) in same command
dapr run --app-id producer --resources-path ./components --dapr-http-port 6001 go run http/producer/main.go 6001
```

### Reference
1. [PubSub](https://docs.dapr.io/developing-applications/building-blocks/pubsub/howto-publish-subscribe/)
1. [Retry / Resiliency Policies](https://docs.dapr.io/operations/resiliency/policies/)
1. [RabbitMQ Configuration](https://docs.dapr.io/reference/components-reference/supported-pubsub/setup-rabbitmq/)
1. [Raw Messages](https://docs.dapr.io/developing-applications/building-blocks/pubsub/pubsub-raw)
1. [Performance](https://docs.dapr.io/operations/performance-and-scalability/perf-service-invocation/)
