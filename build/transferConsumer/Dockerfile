FROM golang:1.18-alpine

WORKDIR /usr/src/consumer

COPY pkg ./pkg
COPY internal ./internal
COPY cmd/transferConsumer/main.go ./

RUN go mod init consumer
RUN go mod tidy

RUN go install consumer

CMD ["consumer"]