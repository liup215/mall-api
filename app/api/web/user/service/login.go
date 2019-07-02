package service

import (
	"errors"
	"fmt"
	"mall/app/api/web/user/model"
)

func (s *Service) Login(username, password string) (*model.Users, error) {
	u, err := s.d.QueryUser(model.UserQuery{Username: username})
	if err != nil {
		return nil, errors.New("账号或密码错误")
	}

	fmt.Println("正确的加密:", u.Password)
	fmt.Println("错误的加密:", s.userHash(password, u.Salt))

	if u.Password != s.userHash(password, u.Salt) {
		return nil, errors.New("账号或密码错误")
	}

	return u, nil
}
