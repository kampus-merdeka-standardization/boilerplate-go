package rest_client_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSentRequest(t *testing.T) {
	client, clientWithTls := setupTest()

	t.Run("Http Get Request Test", func(t *testing.T) {
		resp, code, err := client.SendRequest(http.MethodGet, "/objects/3", nil, nil)
		require.Nil(t, err)
		assert.Equal(t, 200, code)
		require.NotNil(t, resp)
	})

	t.Run("Http Post Request Test", func(t *testing.T) {
		type bodyData struct {
			Year         int     `json:"year"`
			Price        float64 `json:"price"`
			CPUModel     string  `json:"CPU model"`
			HardDiskSize string  `json:"Hard disk size"`
		}
		type bodyRequest struct {
			Name string   `json:"name"`
			Data bodyData `json:"data"`
		}

		reqBody := bodyRequest{
			Name: "Lenovo Thinkpad 5",
			Data: bodyData{
				Year:         2021,
				Price:        599.99,
				CPUModel:     "Intel Core i5 Gen 10",
				HardDiskSize: "1 TB",
			},
		}

		resp, code, err := client.SendRequest(http.MethodPost, "/objects", reqBody, map[string]string{
			"content-type": "application/json",
		})
		require.Nil(t, err)
		assert.Equal(t, 200, code)
		require.NotNil(t, resp)

		type bodyResponse struct {
			Id        string   `json:"id"`
			Name      string   `json:"name"`
			Data      bodyData `json:"data"`
			CreatedAt string   `json:"createdAt"`
		}

		respBody := bodyResponse{}
		err = json.Unmarshal(resp, &respBody)
		require.Nil(t, err)

		assert.NotEmpty(t, respBody.Id)
		assert.NotEmpty(t, respBody.CreatedAt)
		assert.Equal(t, reqBody.Name, respBody.Name)
		assert.Equal(t, reqBody.Data, respBody.Data)
	})

	t.Run("Https Get Request Test", func(t *testing.T) {
		resp, code, err := clientWithTls.SendRequest(http.MethodGet, "/objects/3", nil, nil)
		require.Nil(t, err)
		assert.Equal(t, 200, code)
		require.NotNil(t, resp)
	})

	t.Run("Https Post Request Test", func(t *testing.T) {
		type bodyData struct {
			Year         int     `json:"year"`
			Price        float64 `json:"price"`
			CPUModel     string  `json:"CPU model"`
			HardDiskSize string  `json:"Hard disk size"`
		}
		type bodyRequest struct {
			Name string   `json:"name"`
			Data bodyData `json:"data"`
		}

		reqBody := bodyRequest{
			Name: "Lenovo Thinkpad 5",
			Data: bodyData{
				Year:         2021,
				Price:        599.99,
				CPUModel:     "Intel Core i5 Gen 10",
				HardDiskSize: "1 TB",
			},
		}

		resp, code, err := clientWithTls.SendRequest(http.MethodPost, "/objects", reqBody, map[string]string{
			"content-type": "application/json",
		})
		require.Nil(t, err)
		assert.Equal(t, 200, code)
		require.NotNil(t, resp)

		type bodyResponse struct {
			Id        string   `json:"id"`
			Name      string   `json:"name"`
			Data      bodyData `json:"data"`
			CreatedAt string   `json:"createdAt"`
		}

		respBody := bodyResponse{}
		err = json.Unmarshal(resp, &respBody)
		require.Nil(t, err)

		assert.NotEmpty(t, respBody.Id)
		assert.NotEmpty(t, respBody.CreatedAt)
		assert.Equal(t, reqBody.Name, respBody.Name)
		assert.Equal(t, reqBody.Data, respBody.Data)
	})
}
