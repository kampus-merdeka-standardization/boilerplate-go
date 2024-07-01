package pinger_resolver

import (
	"context"
	"errors"

	errorPkg "simple-golang-app/pkg/error"
)

func NewPing(ctx context.Context, message string) (*PingResolver, error) {
	if message == "" {
		return nil, errorPkg.NewBadRequest(errors.New("message argument in ping query args is empty"), "Message is empty")
	}

	return &PingResolver{
		message: message,
	}, nil
}
