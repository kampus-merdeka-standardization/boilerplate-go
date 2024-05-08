package rest_client_test

import (
	"path/filepath"
	rest_client "rest-client-example/internal/modules/rest-client"
	mock_http "rest-client-example/test/mock/http"
	"sync"
)

func setupTest() (client, clientWithTls rest_client.RestClient) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		mock_http.StartServer()
	}()

	testDir := getTestDir()
	certFilePath := filepath.Join(testDir, "..", "..", "..", "test", "cert", "mock-client.crt")
	KeyFilePath := filepath.Join(testDir, "..", "..", "..", "test", "cert", "mock-client.key")
	CaCertFilePath := filepath.Join(testDir, "..", "..", "..", "test", "cert", "mock-client.pem")

	client, _ = rest_client.NewRestClient("http://localhost:8080")
	clientWithTls, _ = rest_client.NewRestClientWithTls("http://localhost:8080", certFilePath, KeyFilePath, CaCertFilePath)

	return client, clientWithTls
}

func getTestDir() string {
	// Get the absolute path of the directory containing this file
	absolutePath, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	return absolutePath
}
