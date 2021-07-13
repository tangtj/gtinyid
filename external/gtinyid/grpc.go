package gtinyid

import (
	"context"
	"github.com/tangtj/gtinyid/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/keepalive"
	"sync"
	"time"
)

type GrpcSegmentService struct {
	conn   *grpc.ClientConn
	locker sync.Mutex

	host    string
	bizInfo base.GrpcBizToken

	client base.GrpcSegmentClient
}

func (s *GrpcSegmentService) _init() error {
	s.conn = s._connect()
	s.client = base.NewGrpcSegmentClient(s.conn)
	return nil
}

func (s *GrpcSegmentService) GetSegment() (*base.Segment, error) {
	segInfo, err := s.client.GetSegment(context.TODO(), &s.bizInfo)
	if err != nil {
		return nil, err
	}
	return base.NewSegment(segInfo.BizType, segInfo.StartId, segInfo.Step, segInfo.Incr, segInfo.Remainder), nil
}

func NewGrpcSegmentService(host string, bizType string, token string) base.SegmentService {
	s := GrpcSegmentService{
		host: host,
		bizInfo: base.GrpcBizToken{
			BizType: bizType,
			Token:   token,
		},
	}
	s._init()
	return &s
}

func (s *GrpcSegmentService) _reconnect() {

	for s.conn.WaitForStateChange(context.TODO(), connectivity.Ready) {

		// 链接停机的时候进行重连
		if s.conn.GetState() == connectivity.Shutdown {
			s.locker.Lock()
			s.conn = s._connect()
			s.locker.Unlock()
		}
		time.Sleep(time.Second)
	}

}

func (s *GrpcSegmentService) _connect() *grpc.ClientConn {
	clientConn, _ := grpc.Dial(s.host, grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             20 * time.Second,
		PermitWithoutStream: true, //即使是不活跃的流也会进行 ping参数发送
	}), grpc.WithReturnConnectionError(), grpc.WithInsecure())
	return clientConn
}
