package text

type Tag string

type Aggregator interface {
	Aggregate(<-chan Tag)
	Done() <-chan struct{}
}

type aggregator struct {
	done chan struct{}
}

func NewAggregator() *aggregator {
	return &aggregator{
		done: make(chan struct{}),
	}
}

func (a aggregator) Aggregate(<-chan Tag) {

}

func (a aggregator) Done() <-chan struct{} {
	return a.done
}
