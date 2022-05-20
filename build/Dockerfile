FROM golang:1.18-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY pkg $GOROOT
COPY internal $GOROOT
RUN go mod download

COPY *.go ./

RUN go build -o /wallet

EXPOSE 8080

CMD ["/wallet"]