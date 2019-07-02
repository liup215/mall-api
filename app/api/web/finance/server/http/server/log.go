package server

import (
	"mall/app/api/web/finance/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func logDetail(c *gin.Context) {
	var q model.LogQuery
	if err := c.Bind(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误," + err.Error(),
		})
		return
	}

	log, err := svr.LogDetail(q)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "获取失败," + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    log,
	})
}

func logList(c *gin.Context) {
	var q model.LogQuery

	if err := c.Bind(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误," + err.Error(),
		})
		return
	}

	var p model.Page
	if err := c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误," + err.Error(),
		})
		return
	}

	list, total, err := svr.LogList(q, p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "获取失败," + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"total": total,
			"list":  list,
			"size":  len(list),
		},
	})
}

func logCreate(c *gin.Context) {
	var log model.EweiShopMemberLog
	if err := c.Bind(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误," + err.Error(),
		})
		return
	}
}
