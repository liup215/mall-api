package server

import (
	"errors"
	"mall/lib/net/http"
	"mall/lib/net/http/middleware/auth"

	"github.com/gin-gonic/gin"
	// mtime "mall/lib/time"
)

type Login struct {
	Uniacid int    `json:"uniacid" form:"uniacid" binding:"required"`
	Openid  string `json:"openid" form:"openid" binding:"required"`
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var login Login

	if err := c.Bind(&login); err != nil {
		return nil, errors.New("登录参数错误")
	}

	m, err := svr.WechatLogin(login.Uniacid, login.Openid)
	if err != nil {
		return nil, err
	}

	return &auth.CurrentUser{
		Id:       m.Id,
		Uniacid:  m.Uniacid,
		Uid:      m.Uid,
		Openid:   m.Openid,
		Nickname: m.Nickname,
		Realname: m.Realname,
		Mobile:   m.Mobile,
		Avatar:   m.Avatar,
		Usertype: auth.USER_TYPE_MEMBER,
	}, nil

}

func Unauthorized(c *gin.Context, code int, message string) {
	http.Response(c, code, message, nil)
}
