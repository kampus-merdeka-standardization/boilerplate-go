package rest_client_test

import (
	"os"
	"path/filepath"
	rest_client "rest-client-example/internal/modules/rest-client"
)

func setupTest() (client, clientWithTls rest_client.RestClient) {
	testDir := getTestDir()
	certFilePath := filepath.Join(testDir, "..", "..", "..", "test", "cert", "mock-client.crt")
	KeyFilePath := filepath.Join(testDir, "..", "..", "..", "test", "cert", "mock-client.key")
	CaCertFilePath := filepath.Join(testDir, "..", "..", "..", "test", "cert", "mock-client.pem")

	_, err := os.Stat(certFilePath)
	if err != nil {
		panic(err)
	}
	_, err = os.Stat(KeyFilePath)
	if err != nil {
		panic(err)
	}
	_, err = os.Stat(CaCertFilePath)
	if err != nil {
		panic(err)
	}

	client = rest_client.NewRestClient("https://api.restful-api.dev")
	clientWithTls = rest_client.NewRestClientWithTls("https://api.restful-api.dev", certFilePath, KeyFilePath, CaCertFilePath)

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
