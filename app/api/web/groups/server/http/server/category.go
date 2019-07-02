package server

import (
	"mall/app/api/web/groups/model"
	"mall/lib/net/http"

	"github.com/gin-gonic/gin"
)

func createCategory(c *gin.Context) {
	var cat model.EweiShopGroupsCategory

	if err := c.Bind(&cat); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	xcat, err := svr.CreateCategory(cat)
	if err != nil {
		http.Response(c, 400, err.Error(), nil)
		return
	}

	http.Response(c, 200, "创建成功", xcat)

}

func updateCategory(c *gin.Context) {
	var cat model.EweiShopGroupsCategory

	if err := c.Bind(&cat); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	xcat, err := svr.UpdateCategory(cat)
	if err != nil {
		http.Response(c, 400, err.Error(), nil)
		return
	}

	http.Response(c, 200, "更新成功", xcat)

}

func deleteCategory(c *gin.Context) {
	type ID struct {
		Id int `json:"id"`
	}

	var id ID
	if err := c.Bind(&id); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	err := svr.DeleteCategory(id.Id)
	if err != nil {
		http.Response(c, 400, err.Error(), nil)
		return
	}

	http.Response(c, 200, "删除成功", nil)
}

func displayorderCategory(c *gin.Context) {
	type OrderParam struct {
		Id    int `json:"id"`
		Order int `json:"order"`
	}

	var p OrderParam
	if err := c.Bind(&p); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	err := svr.DisplayorderCategory(p.Id, p.Order)
	if err != nil {
		http.Response(c, 400, err.Error(), nil)
		return
	}

	http.Response(c, 200, "更新成功", nil)
}

func enableCategory(c *gin.Context) {
	type IDS struct {
		Ids []int `json:"ids"`
	}

	var p IDS
	if err := c.Bind(&p); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	err := svr.EnableCategory(p.Ids)
	if err != nil {
		http.Response(c, 400, err.Error(), nil)
		return
	}

	http.Response(c, 200, "更新成功", nil)
}
