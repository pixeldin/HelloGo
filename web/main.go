package main

import (
	"HelloGo/web/common"
	"HelloGo/web/controller"
	"fmt"
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
	session := common.GetSession(c, "pvpMgr")
	if session == nil {
		c.Abort()
		common.Logging("Create session failed.")
		return
	}
	h := gin.H{}
	user := session.Values["user"]
	if user == nil {
		//if unauthorized
		fmt.Println("Nil user from session, redirect to login")
		c.HTML(http.StatusFound, "auth/login", h)
	} else {
		fmt.Println("=====================Redirect with old session=====================")
		//jump to main page
		c.Redirect(http.StatusMovedPermanently, "manager/index")
	}
}
