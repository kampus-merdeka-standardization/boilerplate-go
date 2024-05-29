package query_pkg

import (
	"context"

	pinger_resolver "github.com/kampus-merdeka-standardization/boilerplate-pinger-app/internal/modules/pinger/handler/graphql"
)

type rootResolver struct {
}

type pingQueryArgs struct {
	Message *string
}

func (r *rootResolver) Ping(ctx context.Context, args pingQueryArgs) (*pinger_resolver.PingResolver, error) {
	pResolver, err := pinger_resolver.NewPing(ctx, *args.Message)
	if err != nil {
		return nil, err
	}

	return pResolver, nil
}

type QueryRootResolver interface {
	Ping(ctx context.Context, args pingQueryArgs) (*pinger_resolver.PingResolver, error)
}
