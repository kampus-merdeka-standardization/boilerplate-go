package mock_graphql_query

func NewRootResolver() QueryRootResolver {
	return &rootResolver{}
}
