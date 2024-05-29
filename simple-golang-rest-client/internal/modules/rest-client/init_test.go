package rest_client_test

import (
	rest_client "rest-client-example/internal/modules/rest-client"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitRestClient(t *testing.T) {
	t.Run("Failed Certificate not found", func(t *testing.T) {
		_, err := rest_client.NewRestClientWithTls("https://api.restful-api.dev", "", "", "")
		require.NotNil(t, err)
	})

}
