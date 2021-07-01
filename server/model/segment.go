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
	MaxId     int64
	Step      int32
	Incr      int32
}

func (i *SegmentId) Next() (int64, SegmentStatus) {
	ret := atomic.AddInt64(&i.CurrentId, i.MaxId)
	if ret > i.MaxId {
		return -1, SegmentStatusOver
	}
	threshold := (i.Step * LoadingRatio) / 100
	if ret > int64(threshold) {
		return ret, SegmentStatusLoading
	}
	return ret, SegmentStatusNormal
}
