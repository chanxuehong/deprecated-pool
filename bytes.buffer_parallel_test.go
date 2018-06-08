package pool

import (
	"bytes"
	"testing"
)

var testBytesBufferPool *BytesBufferPool

func init() {
	new := func() *bytes.Buffer {
		return bytes.NewBuffer(make([]byte, 128))
	}
	testBytesBufferPool = NewBytesBufferPool(1e6, new)
}

func _TestBytesBufferPool_Get_Put(t *testing.T) {
	buf := testBytesBufferPool.Get()
	buf.Reset()
	buf.WriteString("test")
	if buf.String() != "test" {
		t.Error("failed")
		testBytesBufferPool.Put(buf)
		return
	}
	testBytesBufferPool.Put(buf)
}

func TestBytesBufferPool_Get_Put_Parallel1(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1e6; i++ {
		_TestBytesBufferPool_Get_Put(t)
	}
}

func TestBytesBufferPool_Get_Put_Parallel2(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1e6; i++ {
		_TestBytesBufferPool_Get_Put(t)
	}
}

func TestBytesBufferPool_Get_Put_Parallel3(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1e6; i++ {
		_TestBytesBufferPool_Get_Put(t)
	}
}

func TestBytesBufferPool_Get_Put_Parallel4(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1e6; i++ {
		_TestBytesBufferPool_Get_Put(t)
	}
}

func TestBytesBufferPool_Get_Put_Parallel5(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1e6; i++ {
		_TestBytesBufferPool_Get_Put(t)
	}
}

func TestBytesBufferPool_Get_Put_Parallel6(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1e6; i++ {
		_TestBytesBufferPool_Get_Put(t)
	}
}

func TestBytesBufferPool_Get_Put_Parallel7(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1e6; i++ {
		_TestBytesBufferPool_Get_Put(t)
	}
}

func TestBytesBufferPool_Get_Put_Parallel8(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1e6; i++ {
		_TestBytesBufferPool_Get_Put(t)
	}
}
