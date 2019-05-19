package service

import (
	"mall/app/service/main/member/model"
)

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
