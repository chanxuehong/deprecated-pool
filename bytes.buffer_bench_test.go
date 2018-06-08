package pool

import (
	"bytes"
	"sync"
	"testing"
)

func BenchmarkBytesBufferPool_GetFromEmptyPool(b *testing.B) {
	var buffer *bytes.Buffer = bytes.NewBuffer([]byte{})
	new := func() *bytes.Buffer {
		return buffer
	}
	pool := NewBytesBufferPool(10, new)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkSyncPool_GetFromEmptyPool(b *testing.B) {
	var buffer interface{} = bytes.NewBuffer([]byte{})
	new := func() interface{} {
		return buffer
	}
	pool := sync.Pool{
		New: new,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkBytesBufferPool_Get_Put_Serial(b *testing.B) {
	var buffer *bytes.Buffer = bytes.NewBuffer([]byte{})
	new := func() *bytes.Buffer {
		return buffer
	}
	pool := NewBytesBufferPool(10, new)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Put(pool.Get())
	}
}

func BenchmarkSyncPool_Get_Put_Serial(b *testing.B) {
	var buffer interface{} = bytes.NewBuffer([]byte{})
	new := func() interface{} {
		return buffer
	}
	pool := sync.Pool{
		New: new,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Put(pool.Get())
	}
}

func BenchmarkBytesBufferPool_Get_Put_Parallel(b *testing.B) {
	var buffer *bytes.Buffer = bytes.NewBuffer([]byte{})
	new := func() *bytes.Buffer {
		return buffer
	}
	pool := NewBytesBufferPool(1024, new)

	b.ReportAllocs()
	b.SetParallelism(64)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x := pool.Get()
			pool.Put(x)
		}
	})
}

func BenchmarkSyncPool_Get_Put_Parallel(b *testing.B) {
	var buffer interface{} = bytes.NewBuffer([]byte{})
	new := func() interface{} {
		return buffer
	}
	pool := sync.Pool{
		New: new,
	}

	b.ReportAllocs()
	b.SetParallelism(64)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x := pool.Get()
			pool.Put(x)
		}
	})
}
