package server

import (
	"fmt"
	"mall/app/api/admin/exam/conf"
	"mall/app/api/admin/exam/service"
	"mall/lib/net/http/middleware/auth"
	"mall/lib/net/http/middleware/cors"

	"github.com/gin-gonic/gin"
)

var svr *service.Service

func Init(c *conf.Config) {
	svr = service.New()

	r := gin.Default()

	router(r, c)

	r.Run(fmt.Sprintf(":%s", c.Http.Port))
}

func router(r *gin.Engine, c *conf.Config) {

	r.Use(cors.Cors())

	middleware := auth.New(c.Auth)
	middleware.IdentityHandler = auth.IdentityHandler
	r.Use(middleware.MiddlewareFunc())

	r.POST("/dashboard", Dashboard)
	// catGroup := r.Group("/category")
	// {
	// 	catGroup.POST("/create", createCategory)
	// 	catGroup.POST("/update", updateCategory)
	// 	catGroup.POST("/delete", deleteCategory)
	// 	catGroup.POST("/displayorder", displayorderCategory)
	// 	catGroup.POST("/enable", enableCategory)
	// }
}
