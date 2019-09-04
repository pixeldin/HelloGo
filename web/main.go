package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
	})
	//r.Static("/views", "/views/index.html")
	r.LoadHTMLGlob("web/views/*") // 目录 下不能有其他文件夹 不会会报错
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "Main website",
		})
	})
	//r.GET("/", func(c *gin.Context) {
	//	c.Redirect(http.StatusFound, "/index")
	//})
	r.Run()
}
