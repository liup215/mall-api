package server

import (
	"mall/app/api/admin/user/model"
	"mall/lib/net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func info(c *gin.Context) {
	user, _ := c.Get("userID")
	http.Response(c, 200, "成功", gin.H{
		"user": user,
	})
}

func userList(c *gin.Context) {
	var query model.UserListQuery

	if err := c.BindJSON(&query); err != nil {
		http.Response(c, 400, "参数解析失败,"+err.Error(), nil)
		return
	}

	list, total, err := svr.UserList(query)

	if err != nil {
		http.Response(c, 400, "数据获取失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "数据获取成功", gin.H{
		"list":  list,
		"size":  len(list),
		"total": total,
		"page":  query.Page,
	})
}

func changeState(c *gin.Context) {
	uidStr := c.Param("uid")

	uid, err := strconv.Atoi(uidStr)

	if err != nil {
		http.Response(c, 400, "参数解析失败,"+err.Error(), nil)
		return
	}

	status, err := svr.ChangeStatus(uid)

	if err != nil {
		http.Response(c, 400, err.Error(), nil)
		return
	} else {
		http.Response(c, 200, "更新成功", gin.H{
			"status": status,
		})
		return
	}
}

func selectUser(c *gin.Context) {

	uidStr := c.Param("uid")

	uid, err := strconv.Atoi(uidStr)

	if err != nil {
		http.Response(c, 400, "参数解析失败,"+err.Error(), nil)
		return
	}

	query := model.UserQuery{
		Uid: uid,
	}

	u, err := svr.QueryUser(query)

	if err != nil {
		http.Response(c, 400, "数据获取失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "数据获取成功", u)
}

func updateUser(c *gin.Context) {
	var param model.UserCreateOrUpdateParam
	if err := c.BindJSON(&param); err != nil {
		http.Response(c, 400, "参数解析失败,"+err.Error(), nil)
		return
	}

	u, err := svr.UpdateUser(param)

	if err != nil {
		http.Response(c, 400, "更新失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "更新成功", u)
}
