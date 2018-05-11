package pool

import (
	"bytes"
	"sync"
)

func NewBytesBufferPool(size int, _new func() *bytes.Buffer) *BytesBufferPool {
	if size <= 0 {
		panic("size must be greater than 0")
	}
	if _new == nil {
		panic("_new cannot be nil")
	}
	return &BytesBufferPool{
		pointer: 0,
		buf:     make([]*bytes.Buffer, size),
		pool: sync.Pool{
			New: func() interface{} { return _new() },
		},
	}
}

type BytesBufferPool struct {
	mu      sync.Mutex // protects following fields
	pointer int
	buf     []*bytes.Buffer

	pool sync.Pool
}

func (p *BytesBufferPool) Put(x *bytes.Buffer) {
	if x == nil {
		return
	}
	buf := p.buf
	p.mu.Lock()
	if pointer := p.pointer; pointer >= 0 && pointer < len(buf) {
		buf[pointer] = x
		p.pointer++
		p.mu.Unlock()
		return
	}
	p.mu.Unlock()
	p.pool.Put(x)
}

func (p *BytesBufferPool) Get() (x *bytes.Buffer) {
	buf := p.buf
	p.mu.Lock()
	if pointer := p.pointer - 1; pointer >= 0 && pointer < len(buf) {
		p.pointer = pointer
		x = buf[pointer]
		p.mu.Unlock()
		return
	}
	p.mu.Unlock()
	return p.pool.Get().(*bytes.Buffer)
}
