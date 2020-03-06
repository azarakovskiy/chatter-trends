package main

import (
	"github.com/taxibeat/hackathon-chatter-trends/internal/app/audio"
	listener "github.com/taxibeat/hackathon-chatter-trends/internal/app/audio/listener"
	"github.com/taxibeat/hackathon-chatter-trends/internal/app/text"
)

func main() {
	dataCh := make(chan listener.Data)
	tagCh := make(chan text.Tag)

	l := listener.NewListener()
	go l.Listen(dataCh)

	recognizer := audio.NewRecognizer()
	go recognizer.Recognize(dataCh, tagCh)

	aggregator := text.NewAggregator()
	go aggregator.Aggregate(tagCh)

	<-l.Done()
	//<-recognizer.Done()
	//<-aggregator.Done()
}
