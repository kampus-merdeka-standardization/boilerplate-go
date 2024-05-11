package pinger_client_test

import (
	mock_graphql "graphql-client/mock/graphql"
	"sync"
)

func setupTest() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		mock_graphql.StartGraphqlServer()
	}()
}
