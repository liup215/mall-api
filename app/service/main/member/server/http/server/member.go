package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func memberInfoByMobile(c *gin.Context) {

	uniacid := c.Param("uniacid")
	uid, err := strconv.Atoi(uniacid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误," + err.Error(),
		})
		return
	}

	mobile := c.Param("mobile")

	member, err := svr.MemberInfoByMobile(mobile, uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "用户获取失败," + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "用户获取成功",
		"data":    member,
	})

}
