package rest_client_test

import (
	"encoding/json"
	"net/http"
	rest_client "rest-client-example/rest-client"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSentRequest(t *testing.T) {
	client := rest_client.NewRestClient("https://api.restful-api.dev")

	t.Run("Get Request Test", func(t *testing.T) {
		resp, code, err := client.SendRequest(http.MethodGet, "/objects/3", nil, nil)
		assert.Nil(t, err)
		assert.Equal(t, 200, code)
		assert.NotNil(t, resp)
	})
	t.Run("Post Request Test", func(t *testing.T) {
		type bodyRequest struct {
			Name string `json:"name"`
			Year int    `json:"year"`
		}

		reqBody := bodyRequest{
			Name: "Toyota Corolla",
			Year: 1998,
		}

		resp, code, err := client.SendRequest(http.MethodPost, "/objects", reqBody, map[string]string{
			"content-type": "application/json",
		})
		assert.Nil(t, err)
		assert.Equal(t, 200, code)
		assert.NotNil(t, resp)

		type bodyResponse struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			Year      int    `json:"year"`
			CreatedAt string `json:"createdAt"`
		}

		respBody := bodyResponse{}
		err = json.Unmarshal(resp, &respBody)
		assert.Nil(t, err)

		assert.NotEqual(t, "", respBody.Id)
		assert.NotEqual(t, "", respBody.CreatedAt)
	})
}
