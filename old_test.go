package main_test

import (
	"io"
	"testing"

	"github.com/cloudfoundry/sonde-go/events"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
	"github.com/mailru/easyjson"
)

var (
	marshaler = jsonpb.Marshaler{}

	oldEnvelope = &events.Envelope{
		Origin:    gogoproto.String("my-origin"),
		EventType: events.Envelope_LogMessage.Enum(),
	}
)

// Benchmark old proto marshalling
func BenchmarkOldProtoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := marshaler.Marshal(io.Discard, oldEnvelope)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// Benchmark old proto easyjson marshalling
func BenchmarkOldProtoEasyMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := easyjson.Marshal(oldEnvelope)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// Benchmark old proto gogo marshalling
func BenchmarkOldProtoGogoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := gogoproto.Marshal(oldEnvelope)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// Benchmark old proto unmarshalling
func BenchmarkOldProtoUnmarshal(b *testing.B) {
	buf := &buffer{}
	err := marshaler.Marshal(buf, oldEnvelope)
	if err != nil {
		b.Fatal("Marshaling error:", err)
	}

	var e events.Envelope
	for i := 0; i < b.N; i++ {
		err := jsonpb.Unmarshal(buf, &e)
		if err != nil {
			b.Fatal("Unmarshaling error:", err)
		}
	}
}

// Benchmark old proto easyjson unmarshalling
func BenchmarkOldProtoEasyUnmarshal(b *testing.B) {
	buf := &buffer{}
	err := marshaler.Marshal(buf, oldEnvelope)
	if err != nil {
		b.Fatal("Marshaling error:", err)
	}

	var e events.Envelope
	for i := 0; i < b.N; i++ {
		err := easyjson.Unmarshal(buf.buf, &e)
		if err != nil {
			b.Fatal("Unmarshaling error:", err)
		}
	}
}

// Benchmark old proto gogo unmarshalling
func BenchmarkOldProtoGogoUnmarshal(b *testing.B) {
	buf, err := gogoproto.Marshal(oldEnvelope)
	if err != nil {
		b.Fatal("Marshaling error:", err)
	}

	var e events.Envelope
	for i := 0; i < b.N; i++ {
		err := gogoproto.Unmarshal(buf, &e)
		if err != nil {
			b.Fatal("Unmarshaling error:", err)
		}
	}
}

type buffer struct {
	buf []byte
}

func (b *buffer) Write(p []byte) (n int, err error) {
	b.buf = make([]byte, len(p))
	return copy(b.buf, p), nil
}

func (b *buffer) Read(p []byte) (n int, err error) {
	n = copy(p, b.buf)
	return n, nil
}
