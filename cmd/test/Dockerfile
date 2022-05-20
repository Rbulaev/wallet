FROM golang:1.18-alpine

WORKDIR /app

COPY pkg ./pkg
COPY internal ./internal
COPY go.mod go.sum wallet.go ./

RUN go mod download && go mod verify


#RUN go build -o wallet
#RUN go build -o /test

EXPOSE 8080

CMD ["123124124124"]