package mock_graphql_pinger_resolver

import (
	"context"
	"errors"
)

func NewPing(ctx context.Context, message string) (*PingResolver, error) {
	if message == "" {
		return nil, errors.New("message argument in ping query args is empty")
	}

	return &PingResolver{
		message: message,
	}, nil
}
