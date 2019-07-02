package model

type UpdateCreditParam struct {
	Uniacid int     `json:"uniacid" form:"uniacid" binding:"required"`
	Openid  string  `json:"openid" form:"openid"`
	Mobile  string  `json:"mobile" form:"mobile" binding:"required"`
	Logno   string  `json:"logno" form:"logno"`
	Money   float64 `json:"money" form:"money"`
	Type    int     `json:"type" form:"type"`
	Title   string  `json:"title" form:"title"`
	Remark  string  `json:"remark" form:"remark"`
}
