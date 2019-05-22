package service

import (
	"errors"
	"mall/app/service/main/finance/model"
)

func (s *Service) PreRecharge(param model.PreRechargeParam) (*model.PreRechargeResponse, error) {
	if param.Money <= 0 {
		return nil, errors.New("无效的充值金额")
	}

	if param.Uniacid == 0 {
		return nil, errors.New("无效的商城id")
	}

	if param.Openid == "" {
		return nil, errors.New("无效的用户openid")
	}

	log := model.EweiShopMemberLog{
		Uniacid: param.Uniacid,
		Openid:  param.Openid,
		Money:   param.Money,
		Title:   "商城佣金兑换",
		Status:  0,
	}

	record, err := s.d.CreateLog(log)

	return &model.PreRechargeResponse{
		Uniacid: record.Uniacid,
		Openid:  record.Openid,
		Logno:   record.Logno,
	}, err
}

func (s *Service) RechargeLogConfirm(param model.RechargeLogConfirmParam) (*model.RechargeLogConfirmResponse, error) {
	if param.Uniacid == 0 {
		return nil, errors.New("无效的商城id")
	}

	if param.Logno == "" {
		return nil, errors.New("无效的logno")
	}

	log, _ := s.d.LogByLogno(param.Logno, param.Uniacid)
	if log.Status == 1 {
		return nil, errors.New("订单已确认，不能多次确认")
	}

	err := s.d.UpdateLogStatus(param.Logno, 1, param.Uniacid)
	if err != nil {
		return nil, err
	}
	return &model.RechargeLogConfirmResponse{
		Uniacid: log.Uniacid,
		Openid:  log.Openid,
		Money:   log.Money,
		Type:    log.Type,
		Logno:   log.Logno,
	}, nil

}
