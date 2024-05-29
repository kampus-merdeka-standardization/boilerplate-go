package query_pkg

func NewRootResolver() QueryRootResolver {
	return &rootResolver{}
}
