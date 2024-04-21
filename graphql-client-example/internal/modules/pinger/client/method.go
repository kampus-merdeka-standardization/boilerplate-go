package pinger_client

import (
	"context"
)

func (pc *pingerClient) GetMessage(msg string) (string, error) {
	ctx := context.Background()

	// Use client...
	q := query{}
	variables := map[string]any{
		"message": msg,
	}

	err := pc.client.Query(ctx, &q, variables)
	if err != nil {
		return "", err
	}

	return q.Ping.Message, nil
}
