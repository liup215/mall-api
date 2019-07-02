package server

import (
	"context"
	"mall/app/api/web/member/api"
	"mall/app/api/web/member/model"
	"mall/lib/net/http"
	"mall/lib/time"
	"strconv"

	"github.com/gin-gonic/gin"
)

func memberInfoByMobile(c *gin.Context) {

	uniacid := c.Param("uniacid")
	uid, err := strconv.ParseInt(uniacid, 10, 64)
	if err != nil {
		http.Response(c, 400, "参数错误，"+err.Error(), nil)
		return
	}

	mobile := c.Param("mobile")

	m, err := svr.MemberInfoByMobile(context.Background(), &api.MemberInfoByMobileRequest{
		Uniacid: uid,
		Mobile:  mobile,
	})
	if err != nil {
		http.Response(c, 400, "用户获取失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "用户获取成功", model.EweiShopMember{
		Uniacid:    int(m.Uniacid),
		Id:         int(m.Id),
		Openid:     m.Openid,
		Realname:   m.Realname,
		Mobile:     m.Mobile,
		Weixin:     m.Weixin,
		Createtime: time.Time(m.Createtime),
		Status:     int(m.Status),
		Nickname:   m.Nickname,
		Gender:     int(m.Gender),
		Avatar:     m.Avatar,
		Province:   m.Province,
		City:       m.City,
		Area:       m.Area,
		Salt:       m.Salt,
	})

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

func checkMember(c *gin.Context) {

}
