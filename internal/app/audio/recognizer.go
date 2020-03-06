package audio

import (
	speech "cloud.google.com/go/speech/apiv1"
	"context"
	"fmt"
	"github.com/azarakovskiy/chatter-trends/internal/app/text"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	"io"
	"io/ioutil"
	"log"
)

type Recognizer interface {
	Recognize(<-chan Data, chan<- text.Tag)
	Done() <-chan struct{}
}

type recognizer struct {
	client *speech.Client
	stream speechpb.Speech_StreamingRecognizeClient
}

func NewRecognizer(ctx context.Context) *recognizer {

	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	stream, err := client.StreamingRecognize(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Send the initial configuration message.
	if err := stream.Send(&speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					Encoding:          speechpb.RecognitionConfig_LINEAR16,
					LanguageCode:      "en-US",
					SampleRateHertz:   44100,
					AudioChannelCount: 2,
				},
			},
		},
	}); err != nil {
		log.Fatal(err)
	}

	return &recognizer{
		client: client,
		stream: stream,
	}
}

type endlessFile struct {
	dataCh <-chan Data
	done   chan struct{}
}

func (e endlessFile) Read(p []byte) (n int, err error) {
	//var d []byte
	select {
	case d := <-e.dataCh:
		p = d
	case <-e.done:
		break
	}

	return 1024, nil
}

func (r recognizer) Recognize(dataCh <-chan Data, tagCh chan<- text.Tag) {
	str := endlessFile{
		dataCh: dataCh,
		done:   make(chan struct{}),
	}

	go r.recognize(str)

	go func() {
		for {
			resp, err := r.stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Cannot stream results: %v", err)
			}
			if err := resp.Error; err != nil {
				log.Fatalf("Could not recognize: %v", err)
			}
			for _, result := range resp.Results {
				fmt.Printf("Result: %+v\n", result)
			}
		}
	}()
}

func (r recognizer) recognize(input io.Reader) {
	// Pipe stdin to the API.
	buf := make([]byte, 1024)
	for {
		n, err := input.Read(buf)
		if n > 0 {
			if err := r.stream.Send(&speechpb.StreamingRecognizeRequest{
				StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
					AudioContent: buf[:n],
				},
			}); err != nil {
				log.Printf("Could not send audio: %v", err)
			}
		}
		if err == io.EOF {
			// Nothing else to pipe, close the stream.
			if err := r.stream.CloseSend(); err != nil {
				log.Fatalf("Could not close stream: %v", err)
			}
			return
		}
		if err != nil {
			log.Printf("Could not read from input: %v", err)
			continue
		}
	}
}

func (r recognizer) Done() <-chan struct{} {
	panic("implement me")
}

func (r recognizer) recognize2(file string) error {
	if true {
		panic("wrong function called")
	}

	ctx := context.Background()

	client, err := speech.NewClient(ctx)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// Send the contents of the audio file with the encoding and
	// and sample rate information to be transcripted.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:          speechpb.RecognitionConfig_LINEAR16,
			LanguageCode:      "en-US",
			SampleRateHertz:   44100,
			AudioChannelCount: 2,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})

	// Print the results.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			//fmt.Fprintf(w, "\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
			print(fmt.Sprintf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence))
		}
	}
	return nil
}
