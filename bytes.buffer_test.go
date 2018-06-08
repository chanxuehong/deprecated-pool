package pool

import (
	"bytes"
	"runtime"
	"testing"
)

func TestBytesBufferPool_Get_Put(t *testing.T) {
	// get from empty pool
	{
		new := func() *bytes.Buffer {
			return bytes.NewBuffer(make([]byte, 128))
		}
		pool := NewBytesBufferPool(2, new)

		buf1 := pool.Get()
		if buf1 == nil || !bytes.Equal(buf1.Bytes(), make([]byte, 128)) {
			t.Error("failed")
			return
		}
		buf2 := pool.Get()
		if buf2 == nil || !bytes.Equal(buf2.Bytes(), make([]byte, 128)) {
			t.Error("failed")
			return
		}
		buf3 := pool.Get()
		if buf3 == nil || !bytes.Equal(buf3.Bytes(), make([]byte, 128)) {
			t.Error("failed")
			return
		}
	}

	// put and get the same *bytes.Buffer
	{
		new := func() *bytes.Buffer {
			return bytes.NewBuffer(make([]byte, 128))
		}
		pool := NewBytesBufferPool(2, new)

		buf1 := pool.Get()

		// put
		buf1.Reset()
		buf1.WriteString("test")
		pool.Put(buf1)

		// get
		buf2 := pool.Get()
		if buf2 != buf1 || buf2.String() != "test" {
			t.Error("failed")
			return
		}
	}

	// put and get the same *bytes.Buffer multiple times
	{
		new := func() *bytes.Buffer {
			return bytes.NewBuffer(make([]byte, 128))
		}
		pool := NewBytesBufferPool(2, new)

		buf1 := pool.Get()

		// put
		buf1.Reset()
		buf1.WriteString("test")
		pool.Put(buf1)

		// get
		buf2 := pool.Get()
		if buf2 != buf1 || buf2.String() != "test" {
			t.Error("failed")
			return
		}

		// put
		buf2.Reset()
		buf2.WriteString("test-buf2")
		pool.Put(buf2)

		// get
		buf3 := pool.Get()
		if buf3 != buf1 || buf3.String() != "test-buf2" {
			t.Error("failed")
			return
		}
	}

	// put and get
	// more than the size of pool without runtime.GC()
	{
		new := func() *bytes.Buffer {
			return bytes.NewBuffer(make([]byte, 128))
		}
		pool := NewBytesBufferPool(2, new)

		buf1 := pool.Get()
		buf2 := pool.Get()
		buf3 := pool.Get()

		buf1.Reset()
		buf1.WriteString("test-buf1")
		pool.Put(buf1)
		buf2.Reset()
		buf2.WriteString("test-buf2")
		pool.Put(buf2)
		buf3.Reset()
		buf3.WriteString("test-buf3")
		pool.Put(buf3)

		buf4 := pool.Get()
		if buf4 != buf2 || buf4.String() != "test-buf2" {
			t.Error("failed")
			return
		}
		buf5 := pool.Get()
		if buf5 != buf1 || buf5.String() != "test-buf1" {
			t.Error("failed")
			return
		}

		buf6 := pool.Get()
		if buf6 != buf3 || buf6.String() != "test-buf3" {
			t.Error("failed")
			return
		}

		// must empty
		buf7 := pool.Get()
		if buf7 == nil || !bytes.Equal(buf7.Bytes(), make([]byte, 128)) {
			t.Error("failed")
			return
		}
	}

	// put and get
	// more than the size of pool with runtime.GC()
	{
		new := func() *bytes.Buffer {
			return bytes.NewBuffer(make([]byte, 128))
		}
		pool := NewBytesBufferPool(2, new)

		buf1 := pool.Get()
		buf2 := pool.Get()
		buf3 := pool.Get()

		buf1.Reset()
		buf1.WriteString("test-buf1")
		pool.Put(buf1)
		buf2.Reset()
		buf2.WriteString("test-buf2")
		pool.Put(buf2)
		buf3.Reset()
		buf3.WriteString("test-buf3")
		pool.Put(buf3)

		runtime.GC()

		buf4 := pool.Get()
		if buf4 != buf2 || buf4.String() != "test-buf2" {
			t.Error("failed")
			return
		}
		buf5 := pool.Get()
		if buf5 != buf1 || buf5.String() != "test-buf1" {
			t.Error("failed")
			return
		}

		// must empty
		buf6 := pool.Get()
		if buf6 == nil || !bytes.Equal(buf6.Bytes(), make([]byte, 128)) {
			t.Error("failed")
			return
		}
		buf7 := pool.Get()
		if buf7 == nil || !bytes.Equal(buf7.Bytes(), make([]byte, 128)) {
			t.Error("failed")
			return
		}
	}
}
