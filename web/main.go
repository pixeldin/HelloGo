package main

import (
	"HelloGo/web/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("sessionkey"))

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
	session, e := store.Get(c.Request, "pvpMgr")
	c.Set("pvpSession", session)
	if e != nil {
		fmt.Println("Get session err:", e)
		return
	}
	h := gin.H{}
	//if got user from session
	user := session.Values["user"]
	if user.(string) == "" {
		fmt.Println("Nil user from session, redirect to login")
		c.HTML(http.StatusFound, "auth/login", h)
	} else {
		//jump to main page
		c.Redirect(http.StatusMovedPermanently, "manager/index")
	}
	//if unauthorized
}
