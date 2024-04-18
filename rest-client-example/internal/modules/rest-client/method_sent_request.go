package rest_client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (rc *restClient) SendRequest(method, path string, body any, headers map[string]string) (respBody []byte, code int, err error) {
	bodyReader := new(bytes.Reader)
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return nil, 0, err
		}

		bodyReader = bytes.NewReader(jsonBytes)
	}

	req, err := http.NewRequest(method, rc.baseURL+path, bodyReader)
	if err != nil {
		return nil, 0, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := rc.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return respBody, resp.StatusCode, nil
}
