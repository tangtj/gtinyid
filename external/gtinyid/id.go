package gtinyid

import "github.com/tangtj/gtinyid/base"

func NewIdGenerator(bizType string, token string, url string) base.IdGenerator {
	var httpSegmentService = NewHttpSegmentService(url, bizType, token)
	return base.NewSegmentIdGenerator(bizType, httpSegmentService)
}

func NewGrpcIdGenerator(bizType string, token string, url string) base.IdGenerator {
	var httpSegmentService = NewGrpcSegmentService(url, bizType, token)
	return base.NewSegmentIdGenerator(bizType, httpSegmentService)
}
