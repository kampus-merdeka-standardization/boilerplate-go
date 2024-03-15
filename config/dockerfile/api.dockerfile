FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./build/api/main ./cmd/api

# Path: config/Dockerfile.backend

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/build/api/main /app/apiApp

ENTRYPOINT [ "/app/apiApp" ] 