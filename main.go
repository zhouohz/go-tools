package main

import (
	"github.com/zhouohz/go-tools/captcha/slide"
	"net/http"
)
import "github.com/gin-gonic/gin"

func main() {
	// 创建Gin路由
	r := gin.Default()
	r.Use(Cors())
	// 创建一个处理GET请求的路由
	r.GET("/get", func(c *gin.Context) {

		captcha, s, _ := slide.GetCaptcha()
		// 构建JSON响应
		data := map[string]string{"bg": captcha, "block": s}
		c.JSON(http.StatusOK, data)
	})

	// 启动Gin服务器
	r.Run(":8080")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "false")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
		//
	}
}
