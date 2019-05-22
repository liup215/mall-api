package service

import (
	"errors"
	"mall/app/service/main/member/model"
)

func (s *Service) QueryMember(q model.MemberQuery) (model.EweiShopMember, error) {
	return s.d.QueryMember(q)
}

func (s *Service) MemberInfoById(id, uniacid int) (model.EweiShopMember, error) {
	q := model.MemberQuery{
		Id: id,
	}

	return s.d.QueryMember(q)
}

func (s *Service) MemberInfoByOpenid(openid string, uniacid int) (model.EweiShopMember, error) {
	q := model.MemberQuery{
		Openid:  openid,
		Uniacid: uniacid,
	}

	return s.d.QueryMember(q)
}

func (s *Service) MemberInfoByMobile(mobile string, uniacid int) (model.EweiShopMember, error) {
	q := model.MemberQuery{
		Uniacid: uniacid,
		Mobile:  mobile,
	}

	return s.d.QueryMember(q)
}

func (s *Service) MemberUpdateMobile(id int, mobile string) error {
	if id == 0 {
		return errors.New("请输入用户id")
	}

	if mobile == "" {
		return errors.New("请输入正确的手机号")
	}
	return s.d.UpdateMobile(id, mobile)
}
