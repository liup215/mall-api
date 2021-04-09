package model

type UserListQuery struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
	Username string `json:"username" form:"username"`
}

type UserQuery struct {
	Uid      int    `json:"uid" form:"uid"`
	Username string `json:"username" form:"username"`
}

type UserCreateOrUpdateParam struct {
	Uid      int    `json:"uid" form:"uid"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Status   int    `json:"status" form:"status"`
}
