package segment

import (
	"github.com/tangtj/gtinyid/base"
	"github.com/tangtj/gtinyid/internal/dao"
)

type DbSegmentService struct {
}

func (d *DbSegmentService) GetSegment(bizType string) (*base.Segment, error) {
	return dao.GetNextSegment(bizType)
}
