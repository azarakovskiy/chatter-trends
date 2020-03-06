package listener

import (
	"bytes"
	"fmt"
	"github.com/gordonklaus/portaudio"
	"log"
)

const sampleRate = 44100
const seconds = 1

type Data *bytes.Buffer

type Listener interface {
	Listen(chan<- Data)
	Done() <-chan struct{}
}

type listener struct {
	done chan struct{}
}

func NewListener() *listener {
	return &listener{
		done: make(chan struct{}),
	}
}

func (l listener) Listen(data chan<- Data) {
	portaudio.Initialize()
	defer portaudio.Terminate()

	buffer := make([]int32, sampleRate*seconds)
	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(buffer), func(buff []float32) {
		fmt.Println(buff)
	})

	if err != nil {
		log.Fatalf("error while opening audio stream: %w", err)
	}

	err = stream.Start()
	if err != nil {
		log.Fatalf("error while starting audio stream: %w", err)
	}

	defer stream.Close()
	<-l.done
}

func (l listener) Done() <-chan struct{} {
	return l.done
}
