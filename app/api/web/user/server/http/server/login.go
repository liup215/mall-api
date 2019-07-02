package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mall/lib/net/http"
	"mall/lib/net/http/middleware/auth"
	// mtime "mall/lib/time"
)

type Login struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var login Login

	if err := c.Bind(&login); err != nil {
		return nil, errors.New("登录参数错误")
	}

	u, err := svr.Login(login.Username, login.Password)
	if err != nil {
		return nil, err
	}

	return &auth.CurrentUser{
		Id:       u.Uid,
		Username: u.Username,
		Status:   u.Status,
		Usertype: auth.USER_TYPE_ADMIN,
	}, nil

}

func Unauthorized(c *gin.Context, code int, message string) {
	http.Response(c, code, message, nil)
}
