package gtinyid

import "github.com/tangtj/gtinyid/base"

func NewIdGenerator(bizType string, token string, url string) base.IdGenerator {
	var httpSegmentService = NewHttpSegmentService(url, bizType, token)
	return base.NewSegmentIdGenerator(bizType, httpSegmentService)
}
