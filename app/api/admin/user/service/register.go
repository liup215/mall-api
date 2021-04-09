package service

import (
	"mall/app/service/main/user/model"
)

func (s *Service) Register(param model.RegisterParam) error {
	return s.userService.Register(param)

}
