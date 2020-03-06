package listener

import (
	"fmt"
	"github.com/gordonklaus/portaudio"
	"io"
	"log"
	"math"
)

const sampleRate = 44100
const seconds = 1

type Data []byte

type Listener interface {
	Listen(chan<- Data)
	Done() <-chan struct{}
}

type listener struct {
	i      int
	chunks []Data
	done   chan struct{}
}

func NewListener() *listener {
	return &listener{
		done: make(chan struct{}),
	}
}

func (l *listener) Listen() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	buffer := make([]int32, sampleRate*seconds)
	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(buffer), func(buff []float32) {
		chunk := []byte{}

		for _, i := range buff {
			n := math.Float32bits(i)
			chunk = append(chunk, byte(n))
		}

		l.chunks = append(l.chunks, chunk)

		fmt.Println(len(l.chunks))
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

func (l *listener) Read(p []byte) (int, error) {
	if l.i >= len(l.chunks) {
		return 0, io.EOF
	}

	copy(p, l.chunks[l.i])
	l.i += 1

	return sampleRate, nil
}

func (l listener) Done() <-chan struct{} {
	return l.done
}
