package pinger_graphql

import (
	"github.com/graphql-go/graphql"
)

func NewField() *graphql.Field {
	ql := &pingerGraphql{}

	pinger := &graphql.Field{
		Type:    graphql.String,
		Resolve: ql.Ping,
	}

	return pinger
}
