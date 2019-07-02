package server

import (
	"fmt"
	"mall/app/api/web/user/conf"
	"mall/app/api/web/user/service"
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

	authMiddleware := auth.New(c.Auth)
	authMiddleware.PayloadFunc = auth.PayloadFunc
	authMiddleware.IdentityHandler = auth.IdentityHandler
	authMiddleware.Authenticator = Authenticator
	authMiddleware.Unauthorized = Unauthorized

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", register)
	r.Use(authMiddleware.MiddlewareFunc())
	r.GET("/info", info)

}
