package server

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"mall/app/api/mobile/conf"
	authHandler "mall/app/api/mobile/server/auth"
	"mall/lib/net/http/middleware/auth"
	"mall/lib/net/http/middleware/cors"
	"time"
)

func Init(c *conf.Config) {
	r := gin.Default()

	initRouter(r)

	r.RUN(":" + c.Port)
}

func initRouter(r *gin.Engine) {
	r.Use(cors.Cors())

	authMiddleware := auth.New(c.Auth)
	authMiddleware.PayloadFunc = authHandler.PayloadFunc
	authMiddleware.IdentityHandler = authHandler.IdentityHandler
	authMiddleware.Authenticator = authHandler.Authenticator
	authMiddleware.Authorizator = authHandler.Authorizator
	authMiddleware.Unauthorized = authHandler.Unauthorized

	r.POST("/login", authMiddleware.LoginHandler)

}
