package rest_client

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

func NewRestClient(baseUrl string) *restClient {
	return &restClient{
		baseURL: baseUrl,
		client:  new(http.Client),
	}
}

func NewRestClientWithTls(baseUrl, certPath, keyPath, CAPath string) *restClient {
	// load tls certificates
	clientTLSCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
		panic(err)
	}

	// Configure the client to trust TLS server certs issued by a CA.
	certPool, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}

	if caCertPEM, err := os.ReadFile(CAPath); err != nil {
		panic(err)
	} else if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
		panic("invalid cert in CA PEM")
	}

	tlsConfig := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientTLSCert},
	}
	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{Transport: tr}

	return &restClient{
		baseURL: baseUrl,
		client:  client,
	}
}
