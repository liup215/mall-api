package model

type RechargeConfirmParam struct {
	Uniacid int    `json:"uniacid" form:"uniacid" bind:"required"`
	Openid  string `json:"openid" form:"openid" bind:"required"`
	Logno   string `json:"logno" form:"logno" bind:"required"`
}
