package model

type RegisterParam struct {
	Username string `json:"username" form:"username" binding:"required"`
	Pwd      string `json:"pwd" form:"pwd" binding:"required"`
}

type LoginParam struct {
	Username string `json:"username" form:"username" binding:"required"`
	Pwd      string `json:"pwd" form:"pwd" binding:"required"`
}

type UpdatePwdParam struct {
	Username string `json:"username" form:"username" binding:"required"`
	Newpwd   string `json:"newpwd" form:"newpwd" binding:"required"`
}
