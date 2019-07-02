package service

import (
	"mall/app/api/web/pay/conf"
	payClient "mall/app/service/main/payment/server/rpc/client"
)

type Service struct {
	pay *payClient.Client
}

func New(c *conf.Config) *Service {
	return &Service{
		pay: payClient.New(c.PaymentClient),
	}
}
