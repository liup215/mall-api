package server

import (
	"fmt"
	"mall/app/api/web/wechat/conf"
	"mall/app/api/web/wechat/service"
	"mall/lib/net/http/middleware/cors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("scratchsessionid", store))

	r.Use(cors.Cors())

	authGroup := r.Group("/auth")
	{
		authGroup.GET("/oauthParam", wechatOauthParam)
		authGroup.GET("/userinfo", userInfo)
	}

}
