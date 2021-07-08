package base

import (
	"sync/atomic"
)

type SegmentStatus int

const (
	SegmentStatusNormal  SegmentStatus = 0
	SegmentStatusOver    SegmentStatus = 1
	SegmentStatusLoading SegmentStatus = 2
)

const LoadingRatio = 70

type SegmentService interface {
	GetSegment(bizType string) (*Segment, error)
}

type BizInfoSegmentService interface {
	GetSegment() (*Segment, error)
}

type Segment struct {
	bizType   string
	currentId int64
	startId   int64
	maxId     int64
	step      int64
	incr      int64
	status    SegmentStatus
}

func (i *Segment) Next() (int64, SegmentStatus) {
	if i.status == SegmentStatusOver {
		return -1, i.status
	}
	ret := atomic.AddInt64(&i.currentId, i.incr)
	if ret > i.maxId {
		i.status = SegmentStatusOver
		return -1, i.status
	}
	threshold := i.startId + (i.step*LoadingRatio)/100
	if ret > threshold {
		i.status = SegmentStatusLoading
		return ret, i.status
	}
	return ret, i.status
}

func (i *Segment) Status() SegmentStatus {
	return i.status
}

func (i *Segment) CurrentId() int64 {
	return i.currentId
}

func (i *Segment) StartId() int64 {
	return i.startId
}

func (i *Segment) MaxId() int64 {
	return i.maxId
}

func (i *Segment) Incr() int64 {
	return i.incr
}

func (i *Segment) BizType() string {
	return i.bizType
}

func (i *Segment) Step() int64 {
	return i.step
}

func NewSegment(bizType string, startId, step, incr int64) *Segment {
	return &Segment{
		bizType:   bizType,
		currentId: startId,
		startId:   startId,
		maxId:     startId + step,
		step:      step,
		incr:      incr,
		status:    SegmentStatusNormal,
	}
}

type SegmentInfo struct {
	BizType string `json:"biz_type"`
	StartId int64  `json:"start_id"`
	Step    int64  `json:"step"`
	Incr    int64  `json:"incr"`
}
