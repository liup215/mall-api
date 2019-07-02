package server

import (
	"log"
	"mall/app/api/web/member/api"
	"mall/app/api/web/member/conf"
	"mall/app/api/web/member/service"
	"net"

	"google.golang.org/grpc"
)

func Init(c *conf.Config) {
	lis, err := net.Listen("tcp", c.GrpcConfig)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ss := service.New(c)
	api.RegisterMemberServiceServer(s, ss)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
