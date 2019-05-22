package server

import (
	"github.com/gin-gonic/gin"
	"mall/lib/net/http"
	"strconv"
)

func memberInfoByMobile(c *gin.Context) {

	uniacid := c.Param("uniacid")
	uid, err := strconv.Atoi(uniacid)
	if err != nil {
		http.Response(c, 400, "参数错误，"+err.Error(), nil)
		return
	}

	mobile := c.Param("mobile")

	member, err := svr.MemberInfoByMobile(mobile, uid)
	if err != nil {
		http.Response(c, 400, "用户获取失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "用户获取成功", member)

}

func memberUpdateMobile(c *gin.Context) {
	type Param struct {
		Id     int    `json:"id" form:"id" binding:"required"`
		Mobile string `json:"mobile" form:"id" binding:"required"`
	}

	var p Param

	if err := c.Bind(&p); err != nil {
		http.Response(c, 400, "参数错误，"+err.Error(), nil)
		return
	}

	if err := svr.MemberUpdateMobile(p.Id, p.Mobile); err != nil {
		http.Response(c, 400, "更新错误，"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "更新成功", nil)

}
