package model

type RegisterParam struct {
	Uniacid int    `json:"uniacid" form:"uniacid" binding:"required"`
	Mobile  string `json:"mobile" form:"mobile" binding:"required"`
	Pwd     string `json:"pwd" form:"pwd" binding:"required"`
}

type LoginParam struct {
	Uniacid int    `json:"uniacid" form:"uniacid" binding:"required"`
	Mobile  string `json:"mobile" form:"mobile" binding:"required"`
	Pwd     string `json:"pwd" form:"pwd" binding:"required"`
}
