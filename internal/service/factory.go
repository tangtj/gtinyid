package service

import "sync"

type IdGeneratorFactory struct {
	generators map[string]IdGenerator
	locker     sync.Locker
}

func NewIdGeneratorFactory() *IdGeneratorFactory {
	return &IdGeneratorFactory{generators: map[string]IdGenerator{}, locker: &sync.Mutex{}}
}

func (c IdGeneratorFactory) Next(bizType string) (int64, error) {
	cor := c.GetGenerator(bizType)
	return cor.Next()
}

func (c IdGeneratorFactory) BatchNext(bizType string, size int) ([]int64, error) {
	cor := c.GetGenerator(bizType)
	return cor.BatchNext(size)
}

func (c IdGeneratorFactory) GetGenerator(bizType string) IdGenerator {
	defer c.locker.Unlock()
	c.locker.Lock()
	gen, ok := c.generators[bizType]
	if ok {
		return gen
	}
	gen = NewSegmentIdGenerator(bizType)
	c.generators[bizType] = gen
	return gen
}
