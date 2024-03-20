.PHONY : gen-proto run-grpc build-grpc

gen-proto:
	protoc --go_out=./pkg/proto --go-grpc_out=./pkg/proto pkg/proto/**/*.proto

run-api:
	go run ./cmd/api/main.go

build-api:
	go build -o out/api/main ./cmd/api/main.go

run-grpc: gen-proto
	go run ./cmd/grpc/main.go

build-grpc : gen-proto
	go build -o out/grpc/main ./cmd/api/main.go

run-graphql :
	go run ./cmd/graphql/main.go

build-graphql :
	go build -o out/graphql/main ./cmd/api/main.go

up:
	docker-compose -f ./configs/docker-compose.yml up -d

up-build:
	docker-compose -f ./configs/docker-compose.yml up --build -d

down :
	docker-compose -f ./config/docker-compose.yml down