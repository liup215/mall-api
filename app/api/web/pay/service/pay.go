package service

import (
	payApi "mall/app/service/main/payment/api"
	"mall/app/service/main/payment/server/rpc/client"
)

func (s *Service) Pay(p client.UnifiedOrderParams) (*payApi.UnifiedOrderReply, error) {
	return s.pay.UnifiedOrder(p)

}

func (s *Service) PayComplete(plid int) error {
	_, err := s.pay.PayComplete(plid)
	return err
}
