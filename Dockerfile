FROM golang:latest as api-builder

WORKDIR  /src/github.com/dinumathai/dapr-pubsub
COPY . /src/github.com/dinumathai/dapr-pubsub

RUN CGO_ENABLED=0 go build -o /go/bin/dapr-lib-consumer github.com/dinumathai/dapr-pubsub/dapr-lib/consumer
RUN CGO_ENABLED=0 go build -o /go/bin/dapr-lib-producer github.com/dinumathai/dapr-pubsub/dapr-lib/producer
RUN CGO_ENABLED=0 go build -o /go/bin/http-consumer github.com/dinumathai/dapr-pubsub/http/consumer
RUN CGO_ENABLED=0 go build -o /go/bin/http-producer github.com/dinumathai/dapr-pubsub/http/producer

FROM alpine:latest

# COPY application to workdir
WORKDIR /
COPY --from=api-builder /go/bin/dapr-lib-consumer /dapr-lib-consumer
COPY --from=api-builder /go/bin/dapr-lib-producer /dapr-lib-producer
COPY --from=api-builder /go/bin/http-consumer /http-consumer
COPY --from=api-builder /go/bin/http-producer /http-producer 

RUN chmod a+x /dapr-lib-consumer
RUN chmod a+x /dapr-lib-producer
RUN chmod a+x /http-consumer
RUN chmod a+x /http-producer

# By default http consumer
CMD ["/http-consumer"]
