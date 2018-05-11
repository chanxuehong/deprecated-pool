package pool

import (
	"bytes"
	"sync"
	"testing"
)

func BenchmarkBytesBufferPool_Get(b *testing.B) {
	_new := func() *bytes.Buffer {
		return bytes.NewBuffer([]byte{})
	}
	pool := NewBytesBufferPool(10, _new)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkSyncPool_Get(b *testing.B) {
	_new := func() interface{} {
		return bytes.NewBuffer([]byte{})
	}
	pool := sync.Pool{
		New: _new,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkBytesBufferPool_Get_Put(b *testing.B) {
	_new := func() *bytes.Buffer {
		return bytes.NewBuffer(make([]byte, 128))
	}
	pool := NewBytesBufferPool(10, _new)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Put(pool.Get())
	}
}

func BenchmarkSyncPool_Get_Put(b *testing.B) {
	_new := func() interface{} {
		return bytes.NewBuffer(make([]byte, 128))
	}
	pool := sync.Pool{
		New: _new,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Put(pool.Get())
	}
}

func BenchmarkBytesBufferPool_Get_Put_Parallel(b *testing.B) {
	_new := func() *bytes.Buffer {
		return bytes.NewBuffer(make([]byte, 128))
	}
	pool := NewBytesBufferPool(1024, _new)

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
	_new := func() interface{} {
		return bytes.NewBuffer(make([]byte, 128))
	}
	pool := sync.Pool{
		New: _new,
	}

	b.ReportAllocs()
	b.SetParallelism(64)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x := pool.Get().(*bytes.Buffer)
			pool.Put(x)
		}
	})
}

func BenchmarkBytesBufferPool_Get_Put_Parallel2(b *testing.B) {
	_new := func() *bytes.Buffer {
		return bytes.NewBuffer(make([]byte, 128))
	}
	pool := NewBytesBufferPool(1024, _new)

	b.ReportAllocs()
	b.SetParallelism(64)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x1 := pool.Get()
			x2 := pool.Get()
			pool.Put(x2)
			pool.Put(x1)
		}
	})
}

func BenchmarkSyncPool_Get_Put_Parallel2(b *testing.B) {
	_new := func() interface{} {
		return bytes.NewBuffer(make([]byte, 128))
	}
	pool := sync.Pool{
		New: _new,
	}

	b.ReportAllocs()
	b.SetParallelism(64)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x1 := pool.Get().(*bytes.Buffer)
			x2 := pool.Get().(*bytes.Buffer)
			pool.Put(x2)
			pool.Put(x1)
		}
	})
}
