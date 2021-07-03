package service

import (
	"github.com/tangtj/gtinyid/server/dao"
	"github.com/tangtj/gtinyid/server/model"
	"sync"
)

var _ IdGenerator = (*SegmentIdGenerator)(nil)

type SegmentIdGenerator struct {
	BizType string

	segment     *model.SegmentId
	nextSegment *model.SegmentId
	currentId   int64

	locker sync.Mutex
}

func NewSegmentIdGenerator(bizType string) IdGenerator {
	g := &SegmentIdGenerator{BizType: bizType}
	g._init()
	return g
}

func (g *SegmentIdGenerator) _init() {
	g._loadCurr()
}

func (d *SegmentIdGenerator) Next() (int64, error) {
	var next int64 = 0
	for true {
		id, status := d.segment.Next()
		switch status {
		case model.SegmentStatusOver:
			d._loadCurr()
		case model.SegmentStatusLoading:
			go d._loadNext()
			fallthrough
		case model.SegmentStatusNormal:
			next = id
			goto out
		}
	}
out:
	return next, nil
}

func (d *SegmentIdGenerator) BatchNext(size int) (ret []int64, err error) {
	ret = make([]int64, size)
	for i := 0; i < size; i++ {
		ret[i], err = d.Next()
		if err != nil {
			return nil, err
		}
	}
	return
}

func (d *SegmentIdGenerator) _loadNext() error {
	if d.nextSegment == nil {
		defer d.locker.Unlock()
		d.locker.Lock()

		if d.nextSegment == nil {

			s, err := dao.GetNextSegment(d.BizType)
			if err == nil {
				d.nextSegment = s
			} else {
				return err
			}
			return nil
		}
	}
	return nil
}

func (d *SegmentIdGenerator) _loadCurr() error {

	//当前 号段为空 || 号段已使用完
	if d.segment == nil || d.segment.Status() == model.SegmentStatusOver {
		//备用号段没了
		if d.nextSegment != nil {
			//上锁之后再check一次
			defer d.locker.Unlock()
			d.locker.Lock()

			if d.segment == nil || d.segment.Status() == model.SegmentStatusOver {
				d.segment = d.nextSegment
				d.nextSegment = nil
			}
		} else {
			defer d.locker.Unlock()
			d.locker.Lock()

			if d.segment == nil || d.segment.Status() == model.SegmentStatusOver {

				s, err := dao.GetNextSegment(d.BizType)
				if err == nil {
					d.segment = s
				} else {
					return err
				}
				return nil
			}
		}
		return nil
	}
	return nil
}
