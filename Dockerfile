FROM golang:1.22.0-alpine AS builder

COPY . /github.com/vakhrushevk/auth-service/
WORKDIR /github.com/vakhrushevk/auth-service/

RUN go mod download
RUN go build -o ./bin/auth-service cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/vakhrushevk/auth-service/bin/auth-service .

CMD ["./auth-service"]