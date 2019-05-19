package http

import (
	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(c *gin.Context, code int, message string, data interface{}) {
	res := JsonResponse{
		Code:    code,
		Message: message,
	}

	if data != nil {
		res.Data = data
	}
	c.JSON(code, res)
	return
}
