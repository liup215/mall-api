package service

import (
	"mall/app/api/web/wechat/conf"
	"mall/app/api/web/wechat/model"
	"mall/lib/strings"
	xtime "mall/lib/time"

	"github.com/chanxuehong/wechat/mp/core"
	mpoauth2 "github.com/chanxuehong/wechat/mp/oauth2"
	"github.com/chanxuehong/wechat/mp/user"
	"github.com/chanxuehong/wechat/oauth2"
)

func (s *Service) GetUserInfo(p model.UserInfoParams) (*model.McMappingFans, error) {
	ac, err := s.d.QueryAccountWechatsOne(model.AccountWechatsQuery{Uniacid: p.Uniacid})
	if err != nil {
		return nil, err
	}

	oauth2Endpoint := mpoauth2.NewEndpoint(ac.Key, ac.Secret)
	oauth2Client := oauth2.Client{
		Endpoint: oauth2Endpoint,
	}

	token, err := oauth2Client.ExchangeToken(p.Code)
	if err != nil {
		return nil, err
	}

	fan, err := s.d.QueryMcMappingFansOne(model.FansQuery{Uniacid: p.Uniacid, Openid: token.OpenId})
	if err == nil {
		return fan, err
	}

	// 根据openid获取用户信息
	wsvr, err := s.GetWechatService(p.Uniacid)
	if err != nil {
		return nil, err
	}

	client := core.NewClient(wsvr, nil)
	userinfo, err := user.Get(client, token.OpenId, "")

	if err != nil {
		return nil, err
	}

	record := model.McMappingFans{
		Acid:       ac.Acid,
		Uniacid:    p.Uniacid,
		Openid:     token.OpenId,
		Nickname:   userinfo.Nickname,
		Salt:       strings.Random(8),
		Updatetime: xtime.Now(),
		Follow:     userinfo.IsSubscriber,
		Followtime: xtime.Time(userinfo.SubscribeTime),
		Unionid:    userinfo.UnionId,
		Avatar:     userinfo.HeadImageURL,
	}

	defaultGroup, _ := s.d.QueryMcGroupsOne(model.McGroupsQuery{Uniacid: p.Uniacid, Isdefault: 1})
	data := model.McMembers{
		Uniacid:        p.Uniacid,
		Email:          strings.Md5(token.OpenId) + "@we7.cc",
		Salt:           strings.Random(8),
		Groupid:        defaultGroup.Groupid,
		Nickname:       userinfo.Nickname,
		Avatar:         userinfo.HeadImageURL,
		Gender:         userinfo.Sex,
		Nationality:    userinfo.Country,
		Resideprovince: userinfo.Province + "省",
		Residecity:     userinfo.City + "市",
	}

	m, err := s.d.CreateMcMembers(data)
	if err == nil {
		record.Uid = m.Uid
	}

	data.Password = strings.Md5(token.OpenId + data.Salt + conf.Conf.AuthKey)

	return s.d.CreateMcMappingFans(record)

}
