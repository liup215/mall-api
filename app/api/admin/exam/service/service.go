package service

import (
	examConf "mall/app/service/main/exam/conf"
	examService "mall/app/service/main/exam/service"
)

type Service struct {
	examService *examService.Service
}

func New() *Service {
	examConf.Init()

	return &Service{
		examService: examService.New(examConf.Conf),
	}

}
