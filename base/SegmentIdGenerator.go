package base

import (
	"sync"
)

var _ IdGenerator = (*SegmentIdGenerator)(nil)

type SegmentIdGenerator struct {
	bizType string

	segment     *Segment
	nextSegment *Segment

	segmentService SegmentService

	locker sync.Mutex
}

func NewSegmentIdGenerator(bizType string, segmentService SegmentService) IdGenerator {
	g := &SegmentIdGenerator{bizType: bizType, segmentService: segmentService}
	g._init()
	return g
}

func (d *SegmentIdGenerator) _init() {
	d._loadCurr()
}

func (d *SegmentIdGenerator) Next() (int64, error) {
	var next int64 = 0
	for true {
		id, status := d.segment.Next()
		switch status {
		case SegmentStatusOver:
			d._loadCurr()
		case SegmentStatusLoading:
			go d._loadNext()
			fallthrough
		case SegmentStatusNormal:
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

			s, err := d.segmentService.GetSegment(d.bizType)
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
	if d.segment == nil || d.segment.Status() == SegmentStatusOver {

		defer d.locker.Unlock()
		d.locker.Lock()

		//double check
		if d.segment == nil || d.segment.Status() == SegmentStatusOver {

			// 备用号段还有
			if d.nextSegment != nil {
				d.segment, d.nextSegment = d.nextSegment, nil
			} else {
				if s, err := d.segmentService.GetSegment(d.bizType); err == nil {
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
