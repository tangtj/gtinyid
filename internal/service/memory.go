package service

import (
	"github.com/tangtj/gtinyid/base"
	"sync"
	"sync/atomic"
)

var _ base.IdGenerator = (*MemoryIdGenerator)(nil)

type MemoryIdGenerator struct {
	lastId int64
	locker sync.Locker
}

func NewMemoryGenerator() base.IdGenerator {
	return &MemoryIdGenerator{
		lastId: 0,
		locker: &sync.Mutex{},
	}
}

func (g *MemoryIdGenerator) Next() (int64, error) {
	return atomic.AddInt64(&g.lastId, 1), nil
}

func (g *MemoryIdGenerator) BatchNext(size int) ([]int64, error) {

	ret := make([]int64, size)

	g.locker.Lock()

	//直接取出这部分
	next := g.lastId
	g.lastId = g.lastId + int64(size)

	g.locker.Unlock()

	for i := 0; i < size; i++ {
		next = next + 1
		ret[i] = next
	}
	return ret, nil
}
