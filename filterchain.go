package filterchain

type Filter interface {
	Apply(t interface{}) (interface{})
}

type FilterChain struct {
	Filters []Filter
}

func (chain FilterChain) Apply(i interface{})  interface{} {
	for _, filter := range chain.Filters {
		i = filter.Apply(i)
	}
	return i
}

func New(filters...Filter) FilterChain {
	return FilterChain{append(([]Filter)(nil), filters...)}
}

type DefaultFilter struct {}
func (filter DefaultFilter) Apply(t interface{}) (interface{}) {
	return t
}
