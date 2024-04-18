package rest_client_test

import (
	"path/filepath"
	rest_client "rest-client-example/internal/modules/rest-client"
)

func setupTest() (client, clientWithTls rest_client.RestClient) {
	testDir := getTestDir()
	certFilePath := filepath.Join(testDir, "..", "..", "..", "test", "cert", "mock-client.crt")
	KeyFilePath := filepath.Join(testDir, "..", "..", "..", "test", "cert", "mock-client.key")
	CaCertFilePath := filepath.Join(testDir, "..", "..", "..", "test", "cert", "mock-client.pem")

	client, _ = rest_client.NewRestClient("https://api.restful-api.dev")
	clientWithTls, _ = rest_client.NewRestClientWithTls("https://api.restful-api.dev", certFilePath, KeyFilePath, CaCertFilePath)

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
