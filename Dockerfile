FROM golang:1.24.5-bookworm AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o server main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

CMD ["./server"]
