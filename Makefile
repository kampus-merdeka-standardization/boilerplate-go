run-api:
	go run ./cmd/api/main.go

run-grpc:
	go run ./cmd/grpc/main.go

gen-proto-pinger:
	protoc --go_out=./pkg/proto --go-grpc_out=./pkg/proto pkg/proto/pinger.proto
