package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"mall/app/service/main/member/api"
	"mall/app/service/main/member/model"
	"mall/lib/strings"
	"mall/lib/utils"

	"github.com/jinzhu/gorm"
)

func (s *Service) MemberCreate(c context.Context, r *api.MemberCreateRequest) (*api.MemberCreateReply, error) {
	member := model.EweiShopMember{}

	if r.Uniacid == 0 {
		return nil, errors.New("无效的应用id")
	}

	if r.Mobile == "" {
		return nil, errors.New("请输入正确的手机号码！")
	}

	if r.Pwd == "" {
		return nil, errors.New("请输入密码！")
	}

	if _, err := s.d.QueryMember(model.MemberQuery{Uniacid: int(r.Uniacid), Mobile: r.Mobile}); err != gorm.ErrRecordNotFound {
		return nil, errors.New("此手机号已经注册，请直接登录")
	}

	member.Mobile = r.Mobile
	member.Uniacid = int(r.Uniacid)
	member.Mobileverify = 1

	for {
		member.Openid = fmt.Sprintf("wap_user_%v_%s_%s", r.Uniacid, member.Mobile, strings.Random(4))
		if _, err := s.d.QueryMember(model.MemberQuery{Uniacid: int(r.Uniacid), Openid: member.Openid}); err == gorm.ErrRecordNotFound {
			break
		}
	}

	member.Nickname = r.Mobile[0:3] + "xxxx" + r.Mobile[7:]
	member.Salt = strings.Random(20)
	member.Pwd = strings.Md5(r.Pwd + member.Salt)

	m, err := s.d.CreateMember(member)
	return &api.MemberCreateReply{
		Uniacid: int64(m.Uniacid),
		Id:      int64(m.Id),
		Openid:  m.Openid,
	}, err
}

func (s *Service) Register(param model.RegisterParam) error {
	member := model.EweiShopMember{}

	if param.Uniacid == 0 {
		return errors.New("无效的应用id")
	}

	if param.Mobile == "" {
		return errors.New("请输入正确的手机号码！")
	}

	if param.Pwd == "" {
		return errors.New("请输入密码！")
	}

	if _, err := s.d.QueryMember(model.MemberQuery{Uniacid: param.Uniacid, Mobile: param.Mobile}); err != gorm.ErrRecordNotFound {
		return errors.New("此手机号已经注册，请直接登录")
	}

	member.Mobile = param.Mobile
	member.Uniacid = param.Uniacid
	member.Mobileverify = 1

	for {
		member.Openid = fmt.Sprintf("wap_user_%v_%s_%s", param.Uniacid, member.Mobile, strings.Random(4))
		if _, err := s.d.QueryMember(model.MemberQuery{Openid: member.Openid}); err == gorm.ErrRecordNotFound {
			break
		}
	}

	member.Nickname = param.Mobile[0:3] + "xxxx" + param.Mobile[7:]
	member.Salt = strings.Random(20)
	member.Pwd = strings.Md5(param.Pwd + member.Salt)

	_, err := s.d.CreateMember(member)
	return err
}

func (s *Service) UserCheckWechat(param model.UserCheckWechatParam) (*model.EweiShopMember, error) {
	member := model.EweiShopMember{}
	if param.Uniacid == 0 {
		return nil, errors.New("无效的应用ID")
	}
	if param.Openid == "" {
		return nil, errors.New("无效的用户OpenID")
	}
	fmt.Println(param.Uniacid, param.Openid)
	if m, err := s.d.QueryMember(model.MemberQuery{Uniacid: param.Uniacid, Openid: param.Openid}); err != gorm.ErrRecordNotFound {
		if err == nil {
			return &m, nil
		} else {
			return nil, err
		}
	}
	if param.Uid == 0 {
		return nil, errors.New("无效的uid")
	}

	if param.Openid == "" {
		return nil, errors.New("无效的openid")
	}

	member.Uniacid = param.Uniacid
	member.Uid = param.Uid
	member.Openid = param.Openid
	member.Nickname = param.Nickname
	member.Salt = strings.Random(20)
	member.Pwd = strings.Md5(param.Openid + member.Salt)
	member.Avatar = param.Avatar

	return s.d.CreateMember(member)
}

func (s *Service) Login(param model.LoginParam) (string, error) {
	q := model.MemberQuery{
		Uniacid: param.Uniacid,
		Mobile:  param.Mobile,
	}

	member, err := s.d.QueryMember(q)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("用户不存在")
		}
		return "", errors.New("用户查询错误," + err.Error())
	}

	if strings.Md5(param.Pwd+member.Salt) != member.Pwd {
		return "", errors.New("用户名或密码错误")
	}

	memberMap := utils.StructToMap(member)
	memberMap["ewei_shopv2_member_hash"] = strings.Md5(member.Pwd + member.Salt)
	data, err := json.Marshal(&memberMap)
	if err != nil {
		return "", errors.New("数据错误，" + err.Error())
	}

	return base64.StdEncoding.EncodeToString(data), nil

}

func (s *Service) MemberUpdatePwd(p *model.UpdatePwdParam) error {

	if p.Uniacid == 0 {
		return errors.New("无效的应用id")
	}

	if p.Openid == "" {
		return errors.New("无效的Openid")
	}

	if p.Newpwd == "" {
		return errors.New("请输入密码！")
	}

	member, err := s.d.QueryMember(model.MemberQuery{Uniacid: p.Uniacid, Openid: p.Openid})

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return errors.New("查无此用户")
		} else {
			return errors.New("用户查询失败")
		}
	}

	fmt.Println(p.Newpwd, member.Salt)

	pwd := strings.Md5(p.Newpwd + member.Salt)
	fmt.Println(pwd)
	p.Newpwd = pwd
	return s.d.UpdatePwd(p)
}
