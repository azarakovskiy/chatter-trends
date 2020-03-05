package main

import (
	"github.com/azarakovskiy/chatter-trends/internal/app/audio"
	"github.com/azarakovskiy/chatter-trends/internal/app/text"
)

func main() {
	dataCh := make(chan audio.Data)
	tagCh := make(chan text.Tag)

	listener := audio.NewListener()
	go listener.Listen(dataCh)

	recognizer := audio.NewRecognizer()
	go recognizer.Recognize(dataCh, tagCh)

	aggregator := text.NewAggregator()
	go aggregator.Aggregate(tagCh)

	<-listener.Done()
	<-recognizer.Done()
	<-aggregator.Done()
}
