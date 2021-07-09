package api

import (
	"context"
	"errors"
	"github.com/tangtj/gtinyid/base"
)

type GrpcRequest struct {
}

func (g *GrpcRequest) GetSegment(ctx context.Context, info *base.GrpcBizToken) (*base.GrpcSegmentInfo, error) {
	if !checkBiz(info) {
		return nil, errors.New("token异常")
	}
	if ret, err := segmentService.GetBizSegment(info.BizType); err != nil {
		return nil, err
	} else {
		info := &base.GrpcSegmentInfo{
			BizType: ret.BizType(),
			StartId: ret.StartId(),
			Step:    ret.Step(),
			Incr:    ret.Incr(),
		}
		return info, nil
	}
}

func checkBiz(bizInfo *base.GrpcBizToken) bool {
	if !tokenService.CanGenerate(bizInfo.BizType, bizInfo.Token) {
		return false
	}
	return true
}
