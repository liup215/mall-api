package model

type RegisterParam struct {
	Uniacid      int    `json:"uniacid" form:"uniacid" binding:"required"`
	Mobile       string `json:"mobile" form:"mobile"`
	Pwd          string `json:"pwd" form:"pwd"`
	Nickname     string `json:"nickname" form:"form"`
	Uid          int    `json:"uid" form:"uid"`
	Openid       string `json:"openid" form:"openid"`
	RegisterType string `json:"register_type" form:"register_type"`
}

type LoginParam struct {
	Uniacid int    `json:"uniacid" form:"uniacid" binding:"required"`
	Mobile  string `json:"mobile" form:"mobile" binding:"required"`
	Pwd     string `json:"pwd" form:"pwd"`
}

type UpdatePwdParam struct {
	Uniacid int    `json:"uniacid" form:"uniacid" binding:"required"`
	Openid  string `json:"openid" form:"openid" binding:"required"`
	Newpwd  string `json:"newpwd" form:"newpwd" binding:"required"`
}

type UserCheckWechatParam struct {
	Uniacid  int    `json:"uniacid" form:"uniacid" binding:"required"`
	Nickname string `json:"nickname" form:"form"`
	Uid      int    `json:"uid" form:"uid"`
	Openid   string `json:"openid" form:"openid"`
	Avatar   string `json:"avatar" form:"avatar"`
}
