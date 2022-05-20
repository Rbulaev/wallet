FROM golang:1.18-alpine

WORKDIR /wallet

COPY pkg /wallet/pkg
COPY internal /wallet/internal
COPY cmd/wallet/go.mod cmd/wallet/go.sum cmd/wallet/main.go ./
#COPY cmd/test/go.mod cmd/test/go.sum cmd/test/main.go ./

RUN go mod download && go mod verify

#RUN go build -o wallet
RUN go build -o /wallet/app

#EXPOSE 8080

CMD ["/wallet/app"]