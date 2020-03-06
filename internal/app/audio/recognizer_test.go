package audio

import (
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	r := NewRecognizer(context.Background())

	dataCh := make(chan Data)

	f, err := os.Open("../../../samples/trump.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r.Recognize(dataCh, nil)

	buf := make([]byte, 1024)

	f.Read(buf)
	dataCh <- buf

	time.Sleep(5 * time.Second)
	assert.Fail(t, "end")

}

func TestShortFiles(t *testing.T) {
	r := NewRecognizer(context.Background())

	r.recognize2("../../../samples/trump.wav")
	assert.Fail(t, "end")

}
