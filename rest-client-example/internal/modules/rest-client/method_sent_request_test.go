package rest_client_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	mock_http "rest-client-example/test/mock/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSentRequest(t *testing.T) {
	client, _ := setupTest()

	t.Run("Successfully Sending Get Request", func(t *testing.T) {
		id := "123"
		respBody, code, err := client.SendRequest(http.MethodGet, fmt.Sprintf("/product/%s", id), nil, nil)
		require.Nil(t, err)

		assert.Equal(t, http.StatusOK, code)

		res := new(mock_http.Response)
		err = json.Unmarshal(respBody, res)
		require.Nil(t, err)

		assert.Equal(t, id, res.ID)
		assert.Equal(t, "Sikat Gigi", res.Name)
		assert.Equal(t, 12500, res.Price)
	})
}
