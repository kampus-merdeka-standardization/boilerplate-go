package pinger_graphql

import "github.com/graphql-go/graphql"

func (ql *pingerGraphql) Ping(p graphql.ResolveParams) (interface{}, error) {
	return "pong", nil
}
