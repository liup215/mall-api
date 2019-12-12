package server

import (
	groupModel "mall/app/service/main/exam/model"
	"mall/lib/net/http"

	"github.com/gin-gonic/gin"
)

func createExamGroup(c *gin.Context) {

	var group groupModel.ExamGroup
	if err := c.Bind(&group); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	a, err := svr.CreateGroup(group)
	if err != nil {
		http.Response(c, 400, "活动添加失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "添加成功", a)
}

func listExamGroup(c *gin.Context) {
	var q groupModel.ExamGroupQuery
	if err := c.Bind(&q); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	list, total, err := svr.ListGroup(q)
	if err != nil {
		http.Response(c, 400, "查询错误,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "查询成功", gin.H{
		"list":  list,
		"total": total,
	})

}
