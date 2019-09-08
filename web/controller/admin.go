package controller

import (
	"github.com/gin-gonic/gin"
	//"github.com/gorilla/sessions"
	"net/http"
)

func RegisterAdminRouter(group *gin.RouterGroup) {
	a := group.Group("/manager")
	{
		a.GET("/index", Index)
	}
}

func Index(ctx *gin.Context) {
	h := gin.H{}
	//TODO:// 拦截器...
	//session, exists := ctx.Get("pvpSession")
	//if !exists {
	//	fmt.Println("Nil session...")
	//	return
	//}
	//if got user from session
	//s := session.(Session)
	//user := session.Values["user"]
	//if user.(string) == "" {
	//	fmt.Println("Nil user from session, redirect to login")
	//	c.HTML(http.StatusFound, "auth/login", h)
	//}
	//if unauthorized
	ctx.HTML(http.StatusFound, "manager/index", h)
}
