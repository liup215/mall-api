package auth

import (
	"github.com/appleboy/gin-jwt"
	"mall/lib/net/http"
	mtime "mall/lib/time"
)

type Login struct {
	Uniacid int    `json:"uniacid" form:"uniacid" binding:"required"`
	Mobile  string `json:"mobile" form:"mobile" binding:"required"`
	Pwd     string `json:"pwd" form:"pwd" binding:"required"`
}

type CurrentMember struct {
	Id         int
	Uniacid    int
	Uid        int
	Openid     string
	Realname   string     `json:"realname,omitempty"`
	Mobile     string     `json:"mobile,omitempty"`
	Weixin     string     `json:"weixin,omitempty"`
	Createtime mtime.Time `json:"createtime,omitempty"`
	Status     int        `json:"status,omitempty"`
	Nickname   string     `json:"nickname,omitempty"`
	Gender     int        `json:"gender,omitempty"`
	Avatar     string     `json:"avatar,omitempty"`
	Province   string     `json:"province,omitempty"`
	City       string     `json:"city,omitempty"`
	Area       string     `json:"area,omitempty"`
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*CurrentMember); ok {
		return jwt.MapClaims{
			"id":         v.Id,
			"uniacid":    v.Uniacid,
			"uid":        v.Uid,
			"openid":     v.Openid,
			"realname":   v.Realname,
			"mobile":     v.Mobile,
			"weixin":     v.Weixin,
			"createtime": v.Createtime,
			"status":     v.Status,
			"nickname":   v.Nickname,
			"gender":     v.Gender,
			"avatar":     v.Avatar,
			"province":   v.Province,
			"City":       v.City,
			"Area":       v.Area,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(ctx *gin.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return &CurrentMember{
		Id:         claims["id"].(int),
		Uniacid:    claims["uniacid"].(int),
		Uid:        claims["uid"].(int),
		Openid:     claims["openid"],
		Realname:   claims["realname"],
		Mobile:     claims["mobile"],
		Weixin:     claims["weixin"],
		Createtime: claims["createtime"].(mtime.Time),
		Status:     claims["status"].(int),
		Nickname:   claims["nickname"],
		Gender:     claims["gender"].(int),
		Avatar:     claims["avatar"],
		Province:   claims["province"],
		City:       claims["city"],
		Area:       claims["area"],
	}
}

func Authenticator(ctx *gin.Context) (interface{}, error) {
	var login Login

	if err := c.Bind(&login); err != nil {
		return nil, errors.New("登录参数错误")
	}

	record, err := authSvr.Auth(login.Uniacid, login.Mobile, login.Pwd)
	if err != nil {
		return nil, err
	}

	return &CurrentMember{
		Id, int(record.Id),
		Uniacid:    int(record.Uniacid),
		Uid:        int(record.Uid),
		Openid:     record.Openid,
		Realname:   record.Realname,
		Mobile:     record.Mobile,
		Weixin:     record.Weixin,
		Createtime: time.Time(record.Createtime),
		Status:     int(record.Status),
		Nickname:   record.Nickname,
		Gender:     int(record.Gender),
		Avatar:     record.Avatar,
		Province:   record.Province,
		City:       record.City,
		Area:       record.Area,
	}, err
}

func Authorizator(data interface{}, ctx *gin.Context) bool {
	return true
}

func Unauthorized(c *gin.Context, code int, message string) {
	http.Response(c, code, message, nil)
}
