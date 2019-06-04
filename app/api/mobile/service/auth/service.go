package auth

import (
	"mall/app/api/mobile/conf"
	mmbClient "mall/app/service/main/member/server/grpc/client"
)

type Service struct {
	mclient *mmbClient.Client
}

func New(c *conf.Config) *Service {
	return &Service{
		mclient: mmbClient.New(c.GrpcConfig.MmbServerAddress),
	}
}
