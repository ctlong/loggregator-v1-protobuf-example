package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/cloudfoundry/sonde-go/events"
	"github.com/golang/protobuf/jsonpb"
)

var (
	port      = flag.Int("port", 8080, "port to send messages to")
	marshaler = jsonpb.Marshaler{}
)

func init() {
	log.SetPrefix("old client: ")
}

func makeEnvelope() *events.Envelope {
	origin := "my-origin"
	return &events.Envelope{
		Origin:    &origin,
		EventType: events.Envelope_LogMessage.Enum(),
	}
}

func main() {
	flag.Parse()

	url := fmt.Sprintf("http://0.0.0.0:%d/envelope", *port)

	e := makeEnvelope()
	b := &bytes.Buffer{}
	err := marshaler.Marshal(b, e)
	if err != nil {
		log.Panicf("error marshaling: %s\n", err)
	}

	resp, err := http.Post(url, "application/json", b)
	if err != nil {
		log.Panicf("error making request: %s\n", err)
	}

	log.Printf("status code: %d\n", resp.StatusCode)

	if resp.Body == nil {
		log.Panicln("error: no body provided")
	}

	var respE events.Envelope
	err = jsonpb.Unmarshal(resp.Body, &respE)
	if err != nil {
		log.Printf("error unmarshaling body: %s\n", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("unmarshaled response envelope: %#v\n", respE)
}
