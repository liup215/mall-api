package server

import (
	"fmt"
	"mall/app/service/main/exam/conf"
	"mall/app/service/main/exam/service"
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

	groupGroup := r.Group("/group")
	{
		groupGroup.POST("/create", createExamGroup)
		groupGroup.POST("/list", listExamGroup)
	}

	syllabusGroup := r.Group("/syllabus")
	{
		syllabusGroup.POST("/create", createExamSyllabus)
		syllabusGroup.POST("/list", listExamSyllabus)
	}
	// middleware := auth.New(c.Auth)
	// middleware.IdentityHandler = auth.IdentityHandler
	// r.Use(middleware.MiddlewareFunc())

	// catGroup := r.Group("/category")
	// {
	// 	catGroup.POST("/create", createCategory)
	// 	catGroup.POST("/update", updateCategory)
	// 	catGroup.POST("/delete", deleteCategory)
	// 	catGroup.POST("/displayorder", displayorderCategory)
	// 	catGroup.POST("/enable", enableCategory)
	// }
}
