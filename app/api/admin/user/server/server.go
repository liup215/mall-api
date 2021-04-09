package server

import (
	"fmt"
	"mall/app/api/admin/user/conf"
	"mall/app/api/admin/user/model"
	"mall/app/api/admin/user/service"
	serviceUserModel "mall/app/service/main/user/model"
	"mall/lib/net/http/middleware/auth"
	"mall/lib/net/http/middleware/cors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var svr *service.Service

func initUser() {

	_, err := svr.QueryUser(model.UserQuery{Username: "amdin"})
	if err == gorm.ErrRecordNotFound {
		admin := serviceUserModel.RegisterParam{
			Username: "admin",
			Pwd:      "123456",
		}

		if err := svr.Register(admin); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func Init(c *conf.Config) {
	svr = service.New()

	initUser()

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
	r.POST("/list", userList)
	r.POST("/changeStatus/:uid", changeState)
	r.POST("/select/:uid", selectUser)
	r.POST("/update", updateUser)

}
