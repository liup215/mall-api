package service

import (
	"mall/app/api/web/member/model"
)

func (s *Service) RechargeComplete(param model.RechargeConfirmParam) error {
	// 充值确认
	/*
		res, err := s.d.RechargeLogConfirm(param.Logno, param.Uniacid)
		if err != nil {
			return err
		}
	*/

	p := model.UpdateCreditParam{
		//		Uniacid: res.Uniacid,
		//		Openid:  res.Openid,
		//		Logno:   res.Logno,
		//		Money:   res.Money,
		//		Type:    res.Type,
	}
	// 更新余额
	return s.d.UpdateCredit(p)
}
