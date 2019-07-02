package service

import (
	"errors"
	"mall/app/api/web/member/model"
)

func (s *Service) WechatLogin(uniacid int, openid string) (*model.EweiShopMember, error) {
	if uniacid == 0 {
		return nil, errors.New("无效的用户id")
	}

	if openid == "" {
		return nil, errors.New("无效的openid")
	}
	m, err := s.d.QueryMemberOne(model.MemberQuery{Uniacid: uniacid, Openid: openid})
	if err != nil {
		return nil, err
	}

	return m, nil
}
