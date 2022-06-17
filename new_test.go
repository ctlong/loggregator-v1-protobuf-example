package main_test

import (
	"test/pkg/events"
	"testing"

	"github.com/mailru/easyjson"
	"google.golang.org/protobuf/proto"
)

var (
	origin      = "my-origin"
	newEnvelope = &events.Envelope{
		Origin:    &origin,
		EventType: events.Envelope_LogMessage.Enum(),
	}
)

// Benchmark new proto marshalling
func BenchmarkNewProtoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(newEnvelope)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// Benchmark new proto easyjson marshalling
func BenchmarkNewProtoEasyMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := easyjson.Marshal(newEnvelope)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// Benchmark new proto unmarshalling
func BenchmarkNewProtoUnmarshal(b *testing.B) {
	buf, err := proto.Marshal(newEnvelope)
	if err != nil {
		b.Fatal("Marshaling error:", err)
	}

	var e events.Envelope
	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(buf, &e)
		if err != nil {
			b.Fatal("Unmarshaling error:", err)
		}
	}
}

// TODO: fix this
// Benchmark new proto easyjson unmarshalling
// func BenchmarkNewProtoEasyUnmarshal(b *testing.B) {
// 	buf, err := proto.Marshal(newEnvelope)
// 	if err != nil {
// 		b.Fatal("Marshaling error:", err)
// 	}

// 	var e events.Envelope
// 	for i := 0; i < b.N; i++ {
// 		err := easyjson.Unmarshal(buf, &e)
// 		if err != nil {
// 			b.Fatal("Unmarshaling error:", err)
// 		}
// 	}
// }
