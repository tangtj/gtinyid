package api

import (
	"github.com/tangtj/gtinyid/base"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

var grpcReq = &GrpcRequest{}

func GrpcEnable(wg *sync.WaitGroup) {
	defer wg.Done()
	server := grpc.NewServer()
	listener, err := net.Listen("tcp", ":7080")
	if err != nil {
		log.Printf("监听本地端口异常: %s", err)
		return
	}
	base.RegisterGrpcSegmentServer(server, grpcReq)
	grpcErr := server.Serve(listener)
	log.Printf("服务异常: %s", grpcErr)
}
