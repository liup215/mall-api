package server

import (
	groupModel "mall/app/api/web/groups/model"
	"mall/lib/net/http"
	"mall/lib/net/http/middleware/auth"

	"github.com/gin-gonic/gin"
)

func createActivity(c *gin.Context) {
	user, exists := auth.GetCurrentUser(c)
	if !exists {
		http.Response(c, 401, "用户未登录", nil)
		return
	}

	var act groupModel.EweiShopGroupsActivity
	if err := c.Bind(&act); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	act.Uid = user.Id
	a, err := svr.CreateActivity(act)
	if err != nil {
		http.Response(c, 400, "活动添加失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "添加成功", a)
}

func updateActivity(c *gin.Context) {
	user, exists := auth.GetCurrentUser(c)
	if !exists {
		http.Response(c, 401, "用户未登录", nil)
		return
	}

	var act groupModel.EweiShopGroupsActivity
	if err := c.Bind(&act); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	a, err := svr.UpdateActivity(user.Id, act)
	if err != nil {
		http.Response(c, 400, "活动添加失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "更新成功", a)
}

func listActivity(c *gin.Context) {
	_, exists := auth.GetCurrentUser(c)
	if !exists {
		http.Response(c, 401, "用户未登录", nil)
		return
	}
	var q groupModel.ActivityQuery
	if err := c.Bind(&q); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	list, total := svr.QueryActivityList(q)

	http.Response(c, 200, "获取成功", gin.H{
		"list":  list,
		"total": total,
	})
}

func detailActivity(c *gin.Context) {
	user, exists := auth.GetCurrentUser(c)
	if !exists {
		http.Response(c, 401, "用户未登录", nil)
		return
	}
	var q groupModel.ActivityQuery
	if err := c.Bind(&q); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	if q.Id == "" {
		http.Response(c, 400, "id为空", nil)
		return
	}

	if user.Usertype != 0 {
		if err := svr.CheckVisitor(user, q.Id); err != nil {
			http.Response(c, 400, "访问用户记录失败"+err.Error(), nil)
		}

		svr.IncrActivityVisits(q.Id)
	}

	act, err := svr.QueryActivityById(q.Id)
	if err != nil {
		http.Response(c, 400, "获取失败,"+err.Error(), nil)
		return
	}

	if user.Usertype == 0 {
		http.Response(c, 200, "获取成功", act)
		return
	}

	visitors, vtotal := svr.QueryActVisitorList(act.Id)
	buyers, btotal := svr.QueryActBuyerList(act.Id)

	http.Response(c, 200, "获取成功", gin.H{
		"act": act,
		"visitors": gin.H{
			"list":  visitors,
			"total": vtotal,
		},
		"buyers": gin.H{
			"list":  buyers,
			"total": btotal,
		},
	})
}
