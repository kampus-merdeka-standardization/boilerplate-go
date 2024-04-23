package rest_client

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
)

func NewRestClient(baseUrl string) (*restClient, error) {
	return &restClient{
		baseURL: baseUrl,
		client:  new(http.Client),
	}, nil
}

func NewRestClientWithTls(baseUrl, certPath, keyPath, CAPath string) (*restClient, error) {
	// load tls certificates
	clientTLSCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	// Configure the client to trust TLS server certs issued by a CA.
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	if caCertPEM, err := os.ReadFile(CAPath); err != nil {
		return nil, err
	} else if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
		return nil, err
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
	}, nil
}
