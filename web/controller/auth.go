package controller

import (
	"HelloGo/web/common"
	"HelloGo/web/constant"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

//Click login in login.html invoke this method
func RegisterAuthRouter(group *gin.RouterGroup) {
	a := group.Group("/")
	{
		a.POST("/login", Login)
	}
}

func Login(c *gin.Context) {
	var userName = c.DefaultPostForm("UserName", "defaultUser")
	var pwd = c.DefaultPostForm("Password", "")

	if userName == "pixeldin" && pwd == "123" {
		//auth pass, save to session and redirect to index
		session := common.GetSession(c, constant.SESSION_GLOBAL)
		//fixme: turn into something else
		session.Values[constant.SESSION_USER] = userName
		session.Options = &sessions.Options{
			Path: "/",
			//session expire time
			MaxAge:   300,
			HttpOnly: true,
		}
		session.Save(c.Request, c.Writer)

		c.Redirect(http.StatusMovedPermanently, "manager/index")
	} else {
		c.JSON(200, gin.H{
			"msg":      "Auth failed",
			"userName": userName,
		})
	}
}
