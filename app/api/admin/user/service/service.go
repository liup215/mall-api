package service

import (
	userConf "mall/app/service/main/user/conf"
	userService "mall/app/service/main/user/service"
)

type Service struct {
	userService *userService.Service
}

func New() *Service {
	userConf.Init()

	return &Service{
		userService: userService.New(userConf.Conf),
	}
}
