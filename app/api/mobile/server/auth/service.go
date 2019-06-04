package auth

import (
	"mall/app/api/mobile/conf"
	authService "mall/app/api/mobile/service/auth"
)

var (
	authSvr *authService.Service
)

func init() {
	authSvr = authService.New(conf.Conf)
}
