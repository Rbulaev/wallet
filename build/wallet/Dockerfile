FROM golang:1.18-alpine

WORKDIR /wallet

COPY pkg /wallet/pkg
COPY internal /wallet/internal
COPY cmd/wallet/main.go ./

RUN go mod init wallet
RUN go mod tidy
RUN go mod download && go mod verify

RUN go build -o /wallet/app

EXPOSE 8080

CMD ["/wallet/app"]