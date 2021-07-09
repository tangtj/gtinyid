package segment

import (
	"github.com/tangtj/gtinyid/base"
	"github.com/tangtj/gtinyid/internal/dao"
)

type DbSegmentService struct {
	bizType string
}

func (d *DbSegmentService) GetSegment() (*base.Segment, error) {
	return dao.GetNextSegment(d.bizType)
}

func (d *DbSegmentService) GetBizSegment(bizType string) (*base.Segment, error) {
	return dao.GetNextSegment(d.bizType)
}

func NewDbSegmentService(bizType string) base.SegmentService {
	return &DbSegmentService{bizType: bizType}
}
