package service

import (
	"errors"
	"mall/app/api/web/user/model"
	"mall/lib/strings"

	"github.com/jinzhu/gorm"
)

func (s *Service) Register(param model.RegisterParam) error {
	u := model.Users{}

	if param.Username == "" {
		return errors.New("请输入正确的用户名！")
	}

	if param.Pwd == "" {
		return errors.New("请输入密码！")
	}

	if _, err := s.d.QueryUser(model.UserQuery{Username: param.Username}); err != gorm.ErrRecordNotFound {
		return errors.New("此用户名已经注册，请直接登录")
	}

	/*
		for {
			u.Openid = fmt.Sprintf("wap_user_%v_%s_%s", param.Uniacid, member.Mobile, strings.Random(4))
			if _, err := s.d.QueryMember(model.MemberQuery{Openid: member.Openid}); err == gorm.ErrRecordNotFound {
				break
			}
		}
	*/

	u.Status = 1
	u.Groupid = 1
	u.Salt = strings.Random(10)
	u.Password = strings.Sha1(param.Pwd + u.Salt)

	// 记录group信息

	// 记录profile信息

	_, err := s.d.CreateUser(u)
	return err

}
