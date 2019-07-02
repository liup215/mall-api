package server

import (
	"fmt"
	"mall/app/api/web/pay/conf"
	"mall/app/api/web/pay/service"
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
	r.POST("/pay", pay)
	r.POST("/notify", notify)
}
