package service

import (
	"mall/app/service/main/user/model"
)

func (s *Service) Login(username, password string) (*model.Users, error) {
	return s.userService.Login(username, password)
}
