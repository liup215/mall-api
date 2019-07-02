package service

import (
	"errors"
	"mall/app/api/web/member/model"
)

func (s *Service) UpdateCredit(p model.UpdateCreditParam) error {
	if p.Money == 0 {
		return errors.New("请输入正确的充值金额")
	}

	if p.Mobile == "" {
		return errors.New("无效的用户手机号")
	}

	if p.Uniacid == 0 {
		return errors.New("无效的应用ID")
	}

	mem, err := s.d.QueryMember(model.MemberQuery{Uniacid: p.Uniacid, Mobile: p.Mobile})
	if err != nil {
		return err
	}

	if mem.Openid == "" {
		return errors.New("根据用户手机获取用户openid失败")
	}

	p.Openid = mem.Openid

	s.d.DeleteLog(&model.EweiShopMemberLogQuery{Openid: mem.Openid, Uniacid: p.Uniacid, Type: 0, Status: 0})

	if p.Title == "" {
		p.Title = "便易圈充值"
	}

	log := model.EweiShopMemberLog{
		Uniacid: p.Uniacid,
		Title:   p.Title,
		Openid:  mem.Openid,
		Money:   p.Money,
		Type:    p.Type,
		Status:  1,
		Remark:  p.Remark,
	}

	_, err = s.d.InsertLog(log)
	if err != nil {
		return err
	}

	return s.d.UpdateCredit(p)
}
