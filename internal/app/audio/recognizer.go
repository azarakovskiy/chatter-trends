package audio

import (
	"fmt"
	"github.com/taxibeat/hackathon-chatter-trends/internal/app/audio/listener"
	"github.com/taxibeat/hackathon-chatter-trends/internal/app/text"
)

type Recognizer interface {
	Recognize(<-chan listener.Data, chan<- text.Tag)
	Done() <-chan struct{}
}

type recognizer struct {
	done chan struct{}
}

func NewRecognizer() *recognizer {
	return &recognizer{
		done: make(chan struct{}),
	}
}

func (r recognizer) Recognize(data <-chan listener.Data, tags chan<- text.Tag) {
	fmt.Println("waiting for data")
	d := <-data
	fmt.Println("received data")

	if d != nil {

	}
}

func (r recognizer) Done() <-chan struct{} {
	return r.done
}
