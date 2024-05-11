package pinger_client_test

import (
	"context"
	"fmt"
	pinger_client "grpc-client-example/internal/modules/pinger/client"
	"grpc-client-example/internal/modules/pinger/pinger_grpc_gen"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestPingMethod(t *testing.T) {
	setupGrpcServer()
	t.Run("Successfully Sending Request", func(t *testing.T) {
		conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
		require.Nil(t, err)

		msgReq := "Hello From Azie"
		pClient := pinger_client.NewPingerHandler(conn)

		res, err := pClient.Ping(context.Background(), &pinger_grpc_gen.PingRequest{
			Message: msgReq,
		})
		require.Nil(t, err)

		assert.Equal(t, fmt.Sprintf("pong! %s", msgReq), res.Message)
	})
}
