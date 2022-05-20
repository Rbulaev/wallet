FROM golang:1.18-alpine as builder

WORKDIR /app

COPY internal pkg ./
COPY go.mod go.sum ./
COPY cmd/wallet .
#COPY cmd/wallet/wallet.go ./wallet.go
#COPY cmd/test ./

#RUN go mod download && go mod verify

RUN go build -o /wallet
#RUN go build -o /test

EXPOSE 8080

CMD ["/wallet"]
#CMD ["/test"]