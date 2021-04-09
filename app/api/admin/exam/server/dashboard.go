package server

import (
	"mall/app/service/main/exam/utils"
	"mall/lib/net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {

	monthDayUserActionValue := []int{}
	monthDayDoExamQuestionValue := []int{}
	monthDayText := utils.GetMonthDay(time.Now())

	http.Response(c, 200, "数据获取成功", gin.H{
		"examPaperCount":              svr.ExamPaperSelectAllCount(),
		"questionCount":               svr.QuestionSelectAllCount(),
		"doExamPaperCount":            svr.ExamPaperAnswerSelectAllCount(),
		"doQuestionCount":             svr.ExamPaperQuestionCustomerAnswerSelectAllCount(),
		"monthDayUserActionValue":     monthDayUserActionValue,
		"monthDayDoExamQuestionValue": monthDayDoExamQuestionValue,
		"monthDayText":                monthDayText,
	})
}
