package server

import (
	syllabusModel "mall/app/service/main/exam/model"
	"mall/lib/net/http"

	"github.com/gin-gonic/gin"
)

func createExamSyllabus(c *gin.Context) {

	var syllabus syllabusModel.ExamSyllabus
	if err := c.Bind(&syllabus); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	a, err := svr.CreateSyllabus(syllabus)
	if err != nil {
		http.Response(c, 400, "活动添加失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "添加成功", a)
}

func listExamSyllabus(c *gin.Context) {
	var q syllabusModel.ExamSyllabusQuery
	if err := c.Bind(&q); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	list, total, err := svr.ListSyllabus(q)
	if err != nil {
		http.Response(c, 400, "查询错误,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "查询成功", gin.H{
		"list":  list,
		"total": total,
	})

}
