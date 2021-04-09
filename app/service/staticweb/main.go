package main

import (
	"github.com/gin-gonic/gin"
	"mall/lib/net/http/middleware/cors"
	mstring "mall/lib/strings"
	"strconv"
	"strings"
)

func main() {
	router := gin.Default()

	router.Use(cors.Cors())

	router.Static("/static", "./static")

	router.GET("/ueditor/configAndUpload", func(c *gin.Context) {
		action := c.Query("action")
		if action == "" {
			c.JSON(400, gin.H{
				"message": "action参数不能为空",
			})
			return
		}

		c.JSON(200, gin.H{
			"imageActionName":     "imgUpload",
			"imageFieldName":      "upFile",
			"imageMaxSize":        2048000,
			"imageAllowFiles":     []string{".png", ".jpg", ".jpeg", ".gif", ".bmp"},
			"imageCompressEnable": true,
			"imageCompressBorder": 1600,
			"imageInsertAlign":    "none",
			"imageUrlPrefix":      "",
			"imagePathFormat":     "",
		})
	})

	// Ueditor上传文件
	router.POST("/ueditor/configAndUpload", func(c *gin.Context) {

		// 获取文件文件
		file, _ := c.FormFile("upFile")

		// 获取文件大小
		sizeString := c.PostForm("size")
		size, _ := strconv.Atoi(sizeString)

		// 获取文件类型
		fileTypeOriginal := c.PostForm("type")
		arr := strings.Split(fileTypeOriginal, "/")
		fileType := arr[len(arr)-1]

		fileName := mstring.Random(10)
		dst := "./static/image/" + fileName

		// 上传文件到指定的路径
		c.SaveUploadedFile(file, dst)

		c.JSON(200, gin.H{
			"original": file.Filename,
			"name":     file.Filename,
			"url":      "http://" + c.Request.Host + "/static/image/" + fileName,
			"size":     size,
			"type":     fileType,
			"state":    "SUCCESS",
		})

	})

	router.Run(":8001")
}
