package main

import (
	"encoding/base64"
	"github.com/ClaudioAlbertin/go-fast-protobuf/test/testpb"
	"google.golang.org/protobuf/proto"
	"sync"
	"testing"
)

const (
	presumedNumFieldReads = 5
)

var (
	payload []byte
	stdPool = sync.Pool{
		New: func() any {
			return &testpb.Example{}
		},
	}
)

func init() {
	base64Payload := "ChBhcmJpdHJhcnlfc3RyaW5nEhBhcmJpdHJhcnlfc3RyaW5nGhBhcmJpdHJhcnlfc3RyaW5nIhBhcmJpdHJhcnlfc3RyaW5nKhBhcmJpdHJhcnlfc3RyaW5nMhBhcmJpdHJhcnlfc3RyaW5nOhBhcmJpdHJhcnlfc3RyaW5nQhBhcmJpdHJhcnlfc3RyaW5nShBhcmJpdHJhcnlfc3RyaW5nUhBhcmJpdHJhcnlfc3RyaW5nWhBhcmJpdHJhcnlfc3RyaW5nYhBhcmJpdHJhcnlfc3RyaW5nahBhcmJpdHJhcnlfc3RyaW5nchBhcmJpdHJhcnlfc3RyaW5nehBhcmJpdHJhcnlfc3RyaW5nggEQYXJiaXRyYXJ5X3N0cmluZ4oBEGFyYml0cmFyeV9zdHJpbmeSARBhcmJpdHJhcnlfc3RyaW5nmgEQYXJiaXRyYXJ5X3N0cmluZ6IBEGFyYml0cmFyeV9zdHJpbmeqARBhcmJpdHJhcnlfc3RyaW5nsgEQYXJiaXRyYXJ5X3N0cmluZ7oBEGFyYml0cmFyeV9zdHJpbmfCARBhcmJpdHJhcnlfc3RyaW5nygEQYXJiaXRyYXJ5X3N0cmluZ9IBEGFyYml0cmFyeV9zdHJpbmfaARBhcmJpdHJhcnlfc3RyaW5n4gEQYXJiaXRyYXJ5X3N0cmluZ+oBEGFyYml0cmFyeV9zdHJpbmfyARBhcmJpdHJhcnlfc3RyaW5n+gEQYXJiaXRyYXJ5X3N0cmluZ4ICEGFyYml0cmFyeV9zdHJpbmeKAhBhcmJpdHJhcnlfc3RyaW5nkgIQYXJiaXRyYXJ5X3N0cmluZ5oCEGFyYml0cmFyeV9zdHJpbmeiAhBhcmJpdHJhcnlfc3RyaW5nqgIQYXJiaXRyYXJ5X3N0cmluZ7ICEGFyYml0cmFyeV9zdHJpbme6AhBhcmJpdHJhcnlfc3RyaW5nwgIQYXJiaXRyYXJ5X3N0cmluZ8oCEGFyYml0cmFyeV9zdHJpbmfSAhBhcmJpdHJhcnlfc3RyaW5n2gIQYXJiaXRyYXJ5X3N0cmluZ+ICEGFyYml0cmFyeV9zdHJpbmfqAhBhcmJpdHJhcnlfc3RyaW5n8gIQYXJiaXRyYXJ5X3N0cmluZ/oCEGFyYml0cmFyeV9zdHJpbmeCAxBhcmJpdHJhcnlfc3RyaW5nigMQYXJiaXRyYXJ5X3N0cmluZ5IDEGFyYml0cmFyeV9zdHJpbmc="
	p, err := base64.StdEncoding.DecodeString(base64Payload)
	if err != nil {
		panic(err)
	}

	payload = p
}

func BenchmarkLazyNoField(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		e := &testpb.LazyExample{}
		_ = e.UnmarshalFrom(payload)
	}
}

func BenchmarkLazyLowField(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		e := &testpb.LazyExample{}
		_ = e.UnmarshalFrom(payload)
		for i := 0; i < presumedNumFieldReads; i++ {
			_ = e.GetField1()
		}
	}
}

func BenchmarkLazyHighField(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		e := &testpb.LazyExample{}
		_ = e.UnmarshalFrom(payload)
		for i := 0; i < presumedNumFieldReads; i++ {
			_ = e.GetField50()
		}
	}
}

func BenchmarkStd(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		e := &testpb.Example{}
		_ = proto.Unmarshal(payload, e)

		for i := 0; i < presumedNumFieldReads; i++ {
			_ = e.GetField50()
		}
	}
}

func BenchmarkVTProto(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		e := &testpb.Example{}
		_ = e.UnmarshalVT(payload)

		for i := 0; i < presumedNumFieldReads; i++ {
			_ = e.GetField50()
		}
	}
}

func BenchmarkStdWithPool(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		e := stdPool.Get().(*testpb.Example)
		_ = proto.Unmarshal(payload, e)

		for i := 0; i < presumedNumFieldReads; i++ {
			_ = e.GetField50()
		}

		e.Reset()
		stdPool.Put(e)
	}
}

func BenchmarkVTProtoWithPool(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		e := stdPool.Get().(*testpb.Example)
		_ = e.UnmarshalVT(payload)

		for i := 0; i < presumedNumFieldReads; i++ {
			_ = e.GetField50()
		}

		e.Reset()
		stdPool.Put(e)
	}
}
