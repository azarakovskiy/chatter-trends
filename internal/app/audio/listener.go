package audio

type Data []byte

type Listener interface {
	Listen(chan<- Data)
	Done() <-chan struct{}
}

type listener struct {
}

func NewListener() *listener {
	return &listener{}
}

func (l listener) Listen(chan<- Data) {
	panic("implement me")
}

func (l listener) Done() <-chan struct{} {
	panic("implement me")
}
