package gtinyid

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/tangtj/gtinyid/base"
)

var client = resty.New()

type HttpSegmentService struct {
	url     string
	bizType string
	token   string
}

func (s *HttpSegmentService) GetSegment() (*base.Segment, error) {
	hc := client.R()

	resp, err := hc.SetHeaders(map[string]string{
		"biz_type": s.bizType,
		"token":    s.token,
	}).Get(s.url + "/segment")
	if err != nil {
		return nil, err
	}

	retStruct := struct {
		base.Ret
		Data *base.SegmentInfo `json:"data"`
	}{}
	if err := json.Unmarshal(resp.Body(), &retStruct); err == nil {
		if retStruct.Code != "0" {
			return nil, errors.New(retStruct.Msg)
		}
		info := retStruct.Data
		return base.NewSegment(info.BizType, info.StartId, info.Step, info.Incr), nil
	} else {
		return nil, err
	}

}

func NewHttpSegmentService(url string, bizType string, token string) base.BizInfoSegmentService {
	return &HttpSegmentService{url: url, bizType: bizType, token: token}
}
