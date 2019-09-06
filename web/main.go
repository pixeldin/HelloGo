package main

import (
	"HelloGo/web/controller"
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
	r.LoadHTMLGlob("views/**/*.html") // 目录 下不能有其他文件夹 不会会报错
	r.Static("/static", "./static")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", HomePage)

	g := r.Group("")
	controller.RegisterAuthRouter(g)
	r.Run()
}

func HomePage(c *gin.Context)  {
	h := gin.H{}
	//if unauthorized
	c.HTML(http.StatusFound, "auth/login", h)
	//main page
}

