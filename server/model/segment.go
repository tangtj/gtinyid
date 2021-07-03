package model

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

type SegmentId struct {
	BizType   string
	CurrentId int64
	StartId   int64
	MaxId     int64
	Step      int64
	Incr      int64
	status    SegmentStatus
}

func (i *SegmentId) Next() (int64, SegmentStatus) {
	ret := atomic.AddInt64(&i.CurrentId, i.Incr)
	if ret > i.MaxId {
		i.status = SegmentStatusOver
		return -1, i.status
	}
	threshold := i.StartId + (i.Step*LoadingRatio)/100
	if ret > threshold {
		i.status = SegmentStatusLoading
		return ret, i.status
	}
	return ret, i.status
}

func (i *SegmentId) Status() SegmentStatus {
	return i.status
}
