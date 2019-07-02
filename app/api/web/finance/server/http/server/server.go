package server

import (
	"fmt"
	"mall/app/api/web/finance/conf"
	"mall/app/api/web/finance/service"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

var svr *service.Service

func Init(c *conf.Config) {
	svr = service.New(c)

	r := gin.Default()

	router(r)

	r.Run(fmt.Sprintf(":%s", c.Http.Port))
}

func router(r *gin.Engine) {
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type, token",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	logGroup := r.Group("/log")
	{
		logGroup.GET("/list", logList)
		logGroup.GET("/detail", logDetail)
	}

	rechargeGroup := r.Group("/recharge")
	{
		rechargeGroup.POST("/prerecharge", prerecharge)
	}
}
