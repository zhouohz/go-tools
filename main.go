package main

import (
	"github.com/zhouohz/go-tools/captcha"
	"github.com/zhouohz/go-tools/captcha/core/click"
	"github.com/zhouohz/go-tools/captcha/core/puzzle"
	"github.com/zhouohz/go-tools/captcha/core/seq"
	"github.com/zhouohz/go-tools/captcha/core/slide"
	"github.com/zhouohz/go-tools/captcha/store"
	"net/http"
)
import "github.com/gin-gonic/gin"

type req struct {
	Type  int    `json:"type"`
	Token string `json:"token"`
	Data  string `json:"data"`
}

type get struct {
	Type int `form:"type"`
}

func main() {
	// 创建Gin路由
	r := gin.Default()
	r.Use(Cors())
	// 创建一个处理GET请求的路由
	factory := captcha.NewCaptchaServiceFactory()

	cache := store.NewMemoryCache()
	factory.RegisterService(captcha.Click, click.New(4, 18, cache))
	factory.RegisterService(captcha.Puzzle, puzzle.New(cache))
	factory.RegisterService(captcha.Seq, seq.New(5, cache))
	factory.RegisterService(captcha.Slide, slide.New(5, cache))
	r.GET("/get", func(c *gin.Context) {
		var g get
		if err := c.ShouldBindQuery(&g); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx := c.Request.Context()
		service := factory.GetService(g.Type)
		generate, err := service.Get(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"bgImage": generate.Bg,
			"token":   generate.Token,
			"front":   generate.Front,
		})
	})

	r.POST("/check", func(c *gin.Context) {
		var g req
		if err := c.ShouldBindJSON(&g); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx := c.Request.Context()
		service := factory.GetService(g.Type)
		err := service.Check(ctx, g.Token, g.Data)
		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": 0000,
				"data": map[string]interface{}{
					"result":   false,
					"validate": "",
					"token":    g.Token,
				},
			})
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 0000,
			"data": map[string]interface{}{
				"result":   true,
				"validate": "",
				"token":    g.Token,
			},
		})

	})

	//r.GET("/click/get", func(c *gin.Context) {
	//
	//	//// 构建JSON响应
	//	//data := map[string]string{"bg": captcha, "block": s}
	//
	//	img, err := cclick.Get()
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, err.Error())
	//		return
	//	}
	//	base64, err := image2.ToBase64(img.Bg.(image.Image))
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, err.Error())
	//		return
	//	}
	//	c.JSON(http.StatusOK, map[string]any{
	//		"bgImage": base64,
	//		"front":   img.Front,
	//		"token":   img.Token,
	//	})
	//})
	//
	//r.POST("/click/check", func(c *gin.Context) {
	//	var js req
	//
	//	if err := c.ShouldBindJSON(&js); err != nil {
	//		c.JSON(http.StatusOK, map[string]interface{}{
	//			"code": 0001,
	//			"data": err.Error(),
	//		})
	//		return
	//	}
	//
	//	err := cclick.Check(js.Token, js.Data)
	//	if err != nil {
	//		c.JSON(http.StatusOK, map[string]interface{}{
	//			"code": 0001,
	//			"data": err.Error(),
	//		})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, map[string]any{
	//		"code": 0000,
	//		"data": map[string]interface{}{
	//			"result":   true,
	//			"validate": "",
	//			"token":    js.Token,
	//		},
	//	})
	//})
	//
	//r.GET("/seq/get", func(c *gin.Context) {
	//
	//	//// 构建JSON响应
	//	//data := map[string]string{"bg": captcha, "block": s}
	//	se := seq.New()
	//
	//	if err := se.LoadWordDict("captcha/core/seq/dict.txt"); err != nil {
	//		panic(err)
	//	}
	//
	//	img, err := se.Get()
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, err.Error())
	//		return
	//	}
	//
	//	base64, err := image2.ToBase64(img.Bg.(image.Image))
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, err.Error())
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, map[string]any{
	//		"bgImage": base64,
	//		"token":   img.Token,
	//		"front":   img.Front,
	//	})
	//})
	//
	//r.GET("/puzzle/get", func(c *gin.Context) {
	//
	//	//// 构建JSON响应
	//	//data := map[string]string{"bg": captcha, "block": s}
	//	img, err := ppuzzle.Get()
	//	if err != nil {
	//		c.JSON(http.StatusOK, map[string]interface{}{
	//			"code": 0001,
	//			"data": err.Error(),
	//		})
	//	}
	//
	//	base64, err := image2.ToBase64(img.Bg.(image.Image))
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, err.Error())
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, map[string]any{
	//		"bgImage": base64,
	//		"token":   img.Token,
	//		"front":   img.Front,
	//	})
	//})
	//
	//r.POST("/puzzle/check", func(c *gin.Context) {
	//	var js req
	//
	//	if err := c.ShouldBindJSON(&js); err != nil {
	//		c.JSON(http.StatusOK, map[string]interface{}{
	//			"code": 0001,
	//			"data": map[string]interface{}{
	//				"result":   false,
	//				"validate": "",
	//				"token":    js.Token,
	//			},
	//		})
	//		return
	//	}
	//
	//	err := ppuzzle.Check(js.Token, js.Data)
	//	if err != nil {
	//		c.JSON(http.StatusOK, map[string]interface{}{
	//			"code": 0000,
	//			"data": map[string]interface{}{
	//				"result":   false,
	//				"validate": "",
	//				"token":    js.Token,
	//			},
	//		})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, map[string]any{
	//		"code": 0000,
	//		"data": map[string]interface{}{
	//			"result":   true,
	//			"validate": "",
	//			"token":    js.Token,
	//		},
	//	})
	//})

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
