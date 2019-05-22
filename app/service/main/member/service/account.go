package service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"mall/app/service/main/member/model"
	"mall/lib/strings"
	"mall/lib/utils"
)

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

	for {
		member.Openid = fmt.Sprintf("wap_user_%v_%s_%s", param.Uniacid, member.Mobile, strings.Random(4))
		if _, err := s.d.QueryMember(model.MemberQuery{Openid: member.Openid}); err == gorm.ErrRecordNotFound {
			break
		}
	}

	member.Nickname = param.Mobile[0:3] + "xxxx" + param.Mobile[7:]
	member.Salt = strings.Random(20)
	member.Pwd = strings.Md5(param.Pwd + member.Salt)

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
