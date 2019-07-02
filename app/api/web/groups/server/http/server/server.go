package server

import (
	"fmt"
	"mall/app/api/web/groups/conf"
	"mall/app/api/web/groups/service"
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
	middleware := auth.New(c.Auth)
	middleware.IdentityHandler = auth.IdentityHandler
	r.Use(middleware.MiddlewareFunc())

	catGroup := r.Group("/category")
	{
		catGroup.POST("/create", createCategory)
		catGroup.POST("/update", updateCategory)
		catGroup.POST("/delete", deleteCategory)
		catGroup.POST("/displayorder", displayorderCategory)
		catGroup.POST("/enable", enableCategory)
	}

	goodsGroup := r.Group("/goods")
	{
		goodsGroup.GET("/list", goodsQueryList)
	}

	activityGroup := r.Group("/activity")
	{
		activityGroup.POST("/create", createActivity)
		activityGroup.POST("/update", updateActivity)
		activityGroup.GET("/list", listActivity)
		activityGroup.GET("/detail", detailActivity)
	}

	visitorGroup := r.Group("/visitor")
	{
		visitorGroup.GET("/share_info", getShareVistorInfo)
	}

	orderGroup := r.Group("/order")
	{
		orderGroup.POST("/confirm", confirmOrder)
		orderGroup.GET("/unpayed", unpayedOrder)
	}

	payGroup := r.Group("/pay")
	{
		payGroup.POST("/prepay", prePay)
	}
}
