package pool

import (
	"bytes"
	"runtime"
	"testing"
)

func TestBytesBufferPool_Get_Put(t *testing.T) {
	_new := func() *bytes.Buffer {
		return bytes.NewBuffer(make([]byte, 128))
	}
	pool := NewBytesBufferPool(2, _new)

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

	{
		buf1.Reset()
		buf1.WriteString("test")
		pool.Put(buf1)

		buf1 = pool.Get()
		if buf1 == nil || buf1.String() != "test" {
			t.Error("failed")
			return
		}
	}

	{
		buf1.Reset()
		buf1.WriteString("test")
		pool.Put(buf1)
		buf2.Reset()
		buf2.WriteString("test")
		pool.Put(buf2)
		buf3.Reset()
		buf3.WriteString("test")
		pool.Put(buf3)

		buf4 := pool.Get()
		if buf4 == nil || buf4.String() != "test" {
			t.Error("failed")
			return
		}
		buf5 := pool.Get()
		if buf5 == nil || buf5.String() != "test" {
			t.Error("failed")
			return
		}
		buf6 := pool.Get()
		if buf6 == nil || buf6.String() != "test" {
			t.Error("failed")
			return
		}
		buf7 := pool.Get()
		if buf7 == nil || !bytes.Equal(buf7.Bytes(), make([]byte, 128)) {
			t.Error("failed")
			return
		}
	}

	{ // again
		buf1.Reset()
		buf1.WriteString("test")
		pool.Put(buf1)
		buf2.Reset()
		buf2.WriteString("test")
		pool.Put(buf2)
		buf3.Reset()
		buf3.WriteString("test")
		pool.Put(buf3)

		buf4 := pool.Get()
		if buf4 == nil || buf4.String() != "test" {
			t.Error("failed")
			return
		}
		buf5 := pool.Get()
		if buf5 == nil || buf5.String() != "test" {
			t.Error("failed")
			return
		}
		buf6 := pool.Get()
		if buf6 == nil || buf6.String() != "test" {
			t.Error("failed")
			return
		}
		buf7 := pool.Get()
		if buf7 == nil || !bytes.Equal(buf7.Bytes(), make([]byte, 128)) {
			t.Error("failed")
			return
		}
	}

	{ // again with runtime.GC()
		buf1.Reset()
		buf1.WriteString("test")
		pool.Put(buf1)
		buf2.Reset()
		buf2.WriteString("test")
		pool.Put(buf2)
		buf3.Reset()
		buf3.WriteString("test")
		pool.Put(buf3)
		runtime.GC()

		buf4 := pool.Get()
		if buf4 == nil || buf4.String() != "test" {
			t.Error("failed")
			return
		}
		buf5 := pool.Get()
		if buf5 == nil || buf5.String() != "test" {
			t.Error("failed")
			return
		}
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

	{ // again
		buf1.Reset()
		buf1.WriteString("test")
		pool.Put(buf1)
		buf2.Reset()
		buf2.WriteString("test")
		pool.Put(buf2)
		buf3.Reset()
		buf3.WriteString("test")
		pool.Put(buf3)

		buf4 := pool.Get()
		if buf4 == nil || buf4.String() != "test" {
			t.Error("failed")
			return
		}
		buf5 := pool.Get()
		if buf5 == nil || buf5.String() != "test" {
			t.Error("failed")
			return
		}
		buf6 := pool.Get()
		if buf6 == nil || buf6.String() != "test" {
			t.Error("failed")
			return
		}
		buf7 := pool.Get()
		if buf7 == nil || !bytes.Equal(buf7.Bytes(), make([]byte, 128)) {
			t.Error("failed")
			return
		}
	}

	{ // again with runtime.GC()
		buf1.Reset()
		buf1.WriteString("test")
		pool.Put(buf1)
		buf2.Reset()
		buf2.WriteString("test")
		pool.Put(buf2)
		buf3.Reset()
		buf3.WriteString("test")
		pool.Put(buf3)
		runtime.GC()

		buf4 := pool.Get()
		if buf4 == nil || buf4.String() != "test" {
			t.Error("failed")
			return
		}
		buf5 := pool.Get()
		if buf5 == nil || buf5.String() != "test" {
			t.Error("failed")
			return
		}
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
