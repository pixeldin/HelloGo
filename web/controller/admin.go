package controller

import (
	"HelloGo/web/common"
	"HelloGo/web/constant"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAdminRouter(group *gin.RouterGroup) {
	a := group.Group("/manager")
	{
		a.GET("/index", Index)
	}
}

func Index(c *gin.Context) {
	//Session 拦截器
	session := common.GetSession(c, constant.SESSION_GLOBAL)
	if session == nil {
		c.Abort()
		common.Logging("Create session failed.")
		return
	}
	h := gin.H{}
	user := session.Values[constant.SESSION_USER]
	if user == nil {
		//if unauthorized
		fmt.Println("Nil user from session, redirect to login")
		c.HTML(http.StatusFound, "auth/login", h)
	} else {
		fmt.Println("=====================Redirect with old session=====================")
		//jump to main page
		c.HTML(http.StatusFound, "manager/index", h)
	}
}
