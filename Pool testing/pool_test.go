package main

import (
	"bytes"
	"sync"
	"testing"
)

//go test -bench="Benchmark" -benchtime=5s -benchmem
var psize int = 20000

var bufferPool1 = sync.Pool{
	New: func() interface{} {
		data := make([]byte, psize)
		return &data
	},
}

var bufferPool2 = sync.Pool{
	New: func() interface{} {
		data := make([]byte, psize)
		return data
	},
}

var bufferPool3 = sync.Pool{
	New: func() interface{} {
		data := bytes.NewBuffer(make([]byte, 0, psize))
		return data
	},
}

var data = make([]byte, psize)

func BenchmarkBufferWithPool1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool1.Get().(*[]byte)
		copy(*buf, data)
		bufferPool1.Put(buf)
	}
}

func BenchmarkBufferWithPool2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool2.Get().([]byte)
		copy(buf, data)
		bufferPool1.Put(buf)
	}
}

func BenchmarkBufferWithoutPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]byte, 20000)
		copy(buf, data)
	}
}

func BenchmarkBufferWithPool3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool3.Get().(*bytes.Buffer)
		// copy(buf, data)
		buf.Write(data)
		buf.Reset()
		bufferPool1.Put(buf)
	}
}
