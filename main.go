package main

import (
	"fmt"
	"github.com/zhouohz/go-tools/captcha/core/click"
	"github.com/zhouohz/go-tools/captcha/core/puzzle"
	"github.com/zhouohz/go-tools/captcha/core/seq"
	"github.com/zhouohz/go-tools/captcha/core/slide"
	image2 "github.com/zhouohz/go-tools/core/image"
	"github.com/zhouohz/go-tools/core/util/random"
	"net/http"
)
import "github.com/gin-gonic/gin"

func main() {
	// 创建Gin路由
	r := gin.Default()
	r.Use(Cors())
	// 创建一个处理GET请求的路由
	newSlide := slide.NewSlide(10)
	r.GET("/get", func(c *gin.Context) {

		//// 构建JSON响应
		//data := map[string]string{"bg": captcha, "block": s}
		generate, err := newSlide.Generate()
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, map[string]any{
			"bgImage":    generate.BgImage,
			"blockImage": generate.BlockImage,
			"key":        generate.Key,
		})
	})

	r.POST("/check", func(c *gin.Context) {

		////// 构建JSON响应
		////data := map[string]string{"bg": captcha, "block": s}
		//generate, err := newSlide.Generate()
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, nil)
		//	return
		//}
		//
		//c.JSON(http.StatusOK, map[string]any{
		//	"bgImage":    generate.BgImage,
		//	"blockImage": generate.BlockImage,
		//	"key":        generate.Key,
		//})
		randInt := random.RandInt(10)
		if randInt > 5 {
			c.JSON(http.StatusOK, "ok")
		} else {
			c.JSON(http.StatusOK, "ok1")
		}

	})

	r.GET("/click/get", func(c *gin.Context) {

		//// 构建JSON响应
		//data := map[string]string{"bg": captcha, "block": s}
		img, get := click.Get()
		base64, err := image2.ToBase64(img)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		fmt.Println(get)
		c.JSON(http.StatusOK, map[string]any{
			"bgImage": base64,
			"letters": get.Letter(),
			//"key":        generate.Key,
		})
	})

	r.GET("/seq/get", func(c *gin.Context) {

		//// 构建JSON响应
		//data := map[string]string{"bg": captcha, "block": s}
		se := seq.New()

		if err := se.LoadWordDict("captcha/core/seq/dict.txt"); err != nil {
			panic(err)
		}

		img, _ := se.Get()
		c.JSON(http.StatusOK, img)
	})

	r.GET("/puzzle/get", func(c *gin.Context) {

		//// 构建JSON响应
		//data := map[string]string{"bg": captcha, "block": s}
		imgs := puzzle.Get()
		strings := make([]string, len(imgs))
		for i := range imgs {
			base64, err := image2.ToBase64(imgs[i])
			if err != nil {
				c.JSON(http.StatusInternalServerError, nil)
				return
			}

			strings[i] = base64
		}
		c.JSON(http.StatusOK, map[string]any{
			"bgImage": strings,
			//"key":        generate.Key,
		})
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
