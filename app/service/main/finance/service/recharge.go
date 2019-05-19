package service

import (
	"errors"
	"mall/app/service/main/finance/model"
)

func (s *Service) RechargeSubmit(param model.RechargeParam) error {
	if param.Money <= 0 {
		return errors.New("无效的充值金额")
	}

	if param.Uniacid == 0 {
		return errors.New("无效的商城id")
	}

	if param.Openid == "" {
		return errors.New("无效的用户openid")
	}

	log := model.EweiShopMemberLog{
		Uniacid: param.Uniacid,
		Openid:  param.Openid,
		Money:   param.Money,
		Title:   "商城佣金兑换",
		Status:  0,
	}

	_, err := s.d.CreateLog(log)
	return err
}
