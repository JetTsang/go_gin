package main

import (
	// "encoding/json"
	"net/http"
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	// 1. 创建一个服务
	server := gin.Default()

	// 设置favicon
	server.Use(favicon.New("./public/favicon.ico"))

	// 设置模版的位置
	server.LoadHTMLGlob("public/*")

	server.GET("/", func(ctx *gin.Context) {
		// 返回模版
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "This data is come from Go background.",
		})
	})

	// 定义静态资源目录
	server.Static("/public", "./public")

	// 端口
	server.Run(":80")

}
