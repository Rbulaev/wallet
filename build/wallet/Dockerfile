FROM golang:1.18-alpine

WORKDIR /usr/src/wallet

COPY pkg ./pkg
COPY internal ./internal
COPY cmd/wallet/main.go ./

RUN go mod init wallet
RUN go mod tidy

RUN go install wallet

EXPOSE 8080

CMD ["wallet"]