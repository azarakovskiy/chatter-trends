package text

type Tag string

type Aggregator interface {
	Aggregate(<-chan Tag)
	Done() <-chan struct{}
}

type aggregator struct {
}

func NewAggregator() *aggregator {
	return &aggregator{}
}

func (a aggregator) Aggregate(<-chan Tag) {
	panic("implement me")
}

func (a aggregator) Done() <-chan struct{} {
	panic("implement me")
}
