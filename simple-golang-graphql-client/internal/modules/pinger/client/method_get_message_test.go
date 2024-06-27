package pinger_client_test

import (
	"fmt"
	pinger_client "simple-golang-graphql-client/internal/modules/pinger/client"
	"testing"

	"github.com/hasura/go-graphql-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMessage(t *testing.T) {
	t.Run("Succesffuly Return Response", func(t *testing.T) {
		client := graphql.NewClient("http://localhost:8082/graphql", nil)
		pingerClient := pinger_client.NewPingerClient(client)

		msgRequest := "Hemlo Dunia!"
		msg, err := pingerClient.GetMessage(msgRequest)
		require.Nil(t, err)

		assert.Equal(t, fmt.Sprintf("Pong! from %s", msgRequest), msg)
	})
}
