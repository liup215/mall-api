package auth

import (
	"fmt"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	gojwt "gopkg.in/dgrijalva/jwt-go.v3"
	"mall/lib/time"
)

type CurrentUser struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Status     int       `json:"status"`
	Uniacid    int       `json:"uniacid,omitempty"`
	Uid        int       `json:"uid,omitempty"`
	Openid     string    `json:"openid,omitempty"`
	Realname   string    `json:"realname,omitempty"`
	Mobile     string    `json:"mobile,omitempty"`
	Weixin     string    `json:"weixin,omitempty"`
	Createtime time.Time `json:"createtime,omitempty"`
	Usertype   int       `json:"usertype"`
}

const (
	USER_TYPE_ADMIN = iota
	USER_TYPE_MEMBER
)

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*CurrentUser); ok {
		return jwt.MapClaims{
			"id":         v.Id,
			"username":   v.Username,
			"status":     v.Status,
			"uniacid":    v.Uniacid,
			"uid":        v.Uid,
			"openid":     v.Openid,
			"realname":   v.Realname,
			"mobile":     v.Mobile,
			"weixin":     v.Weixin,
			"createtime": v.Createtime,
			"usertype":   v.Usertype,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(claims gojwt.MapClaims) interface{} {
	fmt.Println(claims)
	return &CurrentUser{
		Id:         int(claims["id"].(float64)),
		Username:   claims["username"].(string),
		Status:     int(claims["status"].(float64)),
		Uniacid:    int(claims["uniacid"].(float64)),
		Uid:        int(claims["uid"].(float64)),
		Openid:     claims["openid"].(string),
		Realname:   claims["realname"].(string),
		Mobile:     claims["mobile"].(string),
		Weixin:     claims["wexin"].(string),
		Createtime: claims["createtime"].(time.Time),
		Usertype:   int(claims["usertype"].(float64)),
	}
}

func GetCurrentUser(c *gin.Context) (*CurrentUser, bool) {
	user, exists := c.Get("userID")
	if !exists {
		return nil, false
	}

	return user.(*CurrentUser), true
}
