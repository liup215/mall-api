package service

import (
	"mall/app/service/main/account/model"
	"mall/lib/strings"
)

func (s *Service) WechatOauthParams(uniacid int) (*model.WechatOauthParamsReply, error) {
	ac, err := s.d.QueryAccountWechatsOne(model.AccountWechatsQuery{Uniacid: uniacid})
	if err != nil {
		return nil, err
	}
	state := strings.Random(20)

	return &model.WechatOauthParamsReply{
		Appid: ac.Key,
		State: state,
	}, nil
}

func (s *Service) WechatUserInfo(p model.WechatUserInfoParams) {}
