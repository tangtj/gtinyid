package service

import (
	"github.com/tangtj/gtinyid/server/model"
	"sync"
	"sync/atomic"
)

var _ IdGenService = (*DatasourceIdGenerator)(nil)

type DatasourceIdGenerator struct {
	segment     *model.SegmentId
	nextSegment *model.SegmentId
	currentId   int64

	locker sync.Locker
}

func (d *DatasourceIdGenerator) Next() (int64, error) {

	seg := *d.segment
	// 这里成员变量 没有使用指针是，每次取都是传递值而非传递指针，无须关心 segment 是不是被替换了
	id := atomic.AddInt64(&d.currentId, int64(seg.Step))
	if id > seg.MaxId {

	}
	return 0, nil
}

func (d *DatasourceIdGenerator) BatchNext(size int) ([]int64, error) {
	panic("implement me")
}

func (d DatasourceIdGenerator) loadNext() error {
	if d.nextSegment == nil {
		if d.nextSegment != nil {
			defer d.locker.Unlock()
			d.locker.Lock()

			d.segment = d.nextSegment
			d.nextSegment = nil
		}
		//TODO ： query segment
		return nil
	}
	return nil
}
