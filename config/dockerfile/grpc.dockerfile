
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./build/grpc/main ./cmd/grpc

# Path: config/Dockerfile.backend

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/build/grpc/main /app/grpcApp

ENTRYPOINT [ "/app/grpcApp" ] 