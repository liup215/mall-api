package service

import (
	"context"
	"errors"
	"mall/app/api/web/member/api"
	"mall/app/api/web/member/model"
)

func (s *Service) MemberInfoByMobile(c context.Context, req *api.MemberInfoByMobileRequest) (*api.MemberInfoReply, error) {
	q := model.MemberQuery{
		Uniacid: int(req.Uniacid),
		Mobile:  req.Mobile,
	}
	m, err := s.d.QueryMember(q)
	if err != nil {
		return nil, err
	}
	return &api.MemberInfoReply{
		Id:         int64(m.Id),
		Uniacid:    int64(m.Uniacid),
		Uid:        int64(m.Uid),
		Openid:     m.Openid,
		Realname:   m.Realname,
		Mobile:     m.Mobile,
		Weixin:     m.Weixin,
		Createtime: int64(m.Createtime),
		Status:     int64(m.Status),
		Nickname:   m.Nickname,
		Gender:     int64(m.Gender),
		Avatar:     m.Avatar,
		Province:   m.Province,
		City:       m.City,
		Area:       m.Area,
		Salt:       m.Salt,
	}, nil
}

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

/*
func (s *Service) MemberInfoByMobile(mobile string, uniacid int) (model.EweiShopMember, error) {
	q := model.MemberQuery{
		Uniacid: uniacid,
		Mobile:  mobile,
	}

	return s.d.QueryMember(q)
}
*/

func (s *Service) MemberUpdateMobile(id int, mobile string) error {
	if id == 0 {
		return errors.New("请输入用户id")
	}

	if mobile == "" {
		return errors.New("请输入正确的手机号")
	}
	return s.d.UpdateMobile(id, mobile)
}
