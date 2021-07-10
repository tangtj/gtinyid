package service

import (
	"github.com/tangtj/gtinyid/base"
	"sync/atomic"
)

var _ base.IdGenerator = (*MemoryIdGenerator)(nil)

type MemoryIdGenerator struct {
	lastId int64
}

func NewMemoryGenerator() base.IdGenerator {
	return &MemoryIdGenerator{
		lastId: 0,
	}
}

func (g *MemoryIdGenerator) Next() (int64, error) {
	return atomic.AddInt64(&g.lastId, 1), nil
}

func (g *MemoryIdGenerator) BatchNext(size int) ([]int64, error) {

	ret := make([]int64, size)
	lastNext := atomic.AddInt64(&g.lastId, int64(size))

	for i, n := 0, lastNext-int64(size); i < size; i++ {
		ret[i] = n
		n++
	}
	return ret, nil
}
