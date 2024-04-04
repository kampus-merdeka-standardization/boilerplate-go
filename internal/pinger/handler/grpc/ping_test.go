package pinger_grpc_test

import (
	"context"
	"log"
	"net"
	"testing"

	pinger "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/grpc"
	pinger_grpc "github.com/kampus-merdeka-standardization/boilerplate-go/internal/pinger/handler/grpc"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/logger"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial()

}

func runServer() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	server := pinger_grpc.NewPingerServer()
	pinger.RegisterPingerServiceServer(s, server)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
}

func TestPingGrpc(t *testing.T) {
	logger.InitLogger("test", "./grpc-test.log")
	runServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pinger.NewPingerServiceClient(conn)
	resp, err := client.Ping(ctx, &pinger.PingRequest{Message: "Test"})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "pong", resp.GetMessage())
}
