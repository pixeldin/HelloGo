package main

import (
	"HelloGo/web/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use()

	r.GET("ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//r.Static("/views", "/views/index.html")
	r.LoadHTMLGlob("views/**/*.html")
	r.Static("/static", "./static")
	r.GET("/index", HomePage)

	g := r.Group("")
	registerRouter(g)
	r.Run()
}

func registerRouter(g *gin.RouterGroup) {
	controller.RegisterAuthRouter(g)
	controller.RegisterAdminRouter(g)
}

func HomePage(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "manager/index")
}
