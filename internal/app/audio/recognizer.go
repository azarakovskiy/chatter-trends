package audio

import "github.com/azarakovskiy/chatter-trends/internal/app/text"

type Recognizer interface {
	Recognize(<-chan Data, chan<- text.Tag)
	Done() <-chan struct{}
}

type recognizer struct {
}

func NewRecognizer() *recognizer {
	return &recognizer{}
}

func (r recognizer) Recognize(<-chan Data, chan<- text.Tag) {
	panic("implement me")
}

func (r recognizer) Done() <-chan struct{} {
	panic("implement me")
}
