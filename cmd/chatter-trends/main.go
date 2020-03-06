package main

import (
	"bufio"
	"fmt"
	listener "github.com/taxibeat/hackathon-chatter-trends/internal/app/audio/listener"
)

func main() {
	//dataCh := make(chan listener.Data)
	//tagCh := make(chan text.Tag)

	l := listener.NewListener()
	go l.Listen()

	scanner := bufio.NewScanner(l)

	for scanner.Scan() {
		fmt.Println(scanner.Bytes())
	}

	//recognizer := audio.NewRecognizer()
	//go recognizer.Recognize(dataCh, tagCh)
	//
	//aggregator := text.NewAggregator()
	//go aggregator.Aggregate(tagCh)

	<-l.Done()
	//<-recognizer.Done()
	//<-aggregator.Done()
}
