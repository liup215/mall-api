package server

import (
	"fmt"
	"mall/app/service/main/member/conf"
	"mall/app/service/main/member/service"
	"mall/lib/net/http/middleware/auth"
	"mall/lib/net/http/middleware/cors"

	"github.com/gin-gonic/gin"
)

var svr *service.Service

func Init(c *conf.Config) {
	svr = service.New(c)
	r := gin.Default()

	router(r, c)

	r.Run(fmt.Sprintf(":%s", c.Http.Port))
}

func router(r *gin.Engine, c *conf.Config) {
	r.Use(cors.Cors())

	r.GET("/info/:uniacid/mobile/:mobile", memberInfoByMobile)
	accountGroup := r.Group("/account")
	{
		accountGroup.POST("/register", register)
		accountGroup.POST("/login", login)
	}

	creditGroup := r.Group("/credit")
	{
		creditGroup.POST("/update", creditUpdate)
	}

	updateGroup := r.Group("/update")
	{
		updateGroup.POST("/mobile", memberUpdateMobile)
		updateGroup.POST("/pwd", memberUpdatePwd)
	}

	wechatGroup := r.Group("/wechat")
	{
		wechatGroup.POST("/user_check", userCheckWechat)
	}

	v2 := r.Group("/v2")
	{
		authMiddleware := auth.New(c.Auth)
		authMiddleware.PayloadFunc = auth.PayloadFunc
		authMiddleware.IdentityHandler = auth.IdentityHandler
		authMiddleware.Authenticator = Authenticator
		authMiddleware.Unauthorized = Unauthorized

		v2.POST("/login", authMiddleware.LoginHandler)
		v2.Use(authMiddleware.MiddlewareFunc())
		{
			v2.GET("/current_user", currentUser)
		}
	}
}
