package service

import (
	"github.com/tangtj/gtinyid/base"
	"github.com/tangtj/gtinyid/internal/service/segment"
	"sync"
)

type IdGeneratorFactory struct {
	generators map[string]base.IdGenerator
	locker     sync.Locker
}

var dbSegmentService base.SegmentService = &segment.DbSegmentService{}

func NewIdGeneratorFactory() *IdGeneratorFactory {
	return &IdGeneratorFactory{generators: map[string]base.IdGenerator{}, locker: &sync.Mutex{}}
}

func (c IdGeneratorFactory) Next(bizType string) (int64, error) {
	cor := c.GetGenerator(bizType)
	return cor.Next()
}

func (c IdGeneratorFactory) BatchNext(bizType string, size int) ([]int64, error) {
	cor := c.GetGenerator(bizType)
	return cor.BatchNext(size)
}

func (c IdGeneratorFactory) GetGenerator(bizType string) base.IdGenerator {
	defer c.locker.Unlock()
	c.locker.Lock()
	gen, ok := c.generators[bizType]
	if ok {
		return gen
	}
	gen = base.NewSegmentIdGenerator(bizType, dbSegmentService)
	c.generators[bizType] = gen
	return gen
}
