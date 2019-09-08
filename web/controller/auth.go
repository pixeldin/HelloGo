package controller

import (
	"HelloGo/web/common"
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

	if userName == "a" && pwd == "123" {
		//auth pass, save to session and redirect to index
		//sion, ok := c.Get("pvpSession")
		session := common.GetSession(c, "pvpMgr")
		//if ok {
		//	session := sion.(*sessions.Session)
			//fixme: turn into something else
			session.Values["user"] = userName
			session.Options = &sessions.Options{
				Path: "/",
				//session expire time
				MaxAge: 300,
				HttpOnly:true,
			}
			session.Save(c.Request, c.Writer)
		//}

		c.Redirect(http.StatusMovedPermanently, "manager/index")
	} else {
		c.JSON(200, gin.H{
			"msg":      "Auth failed",
			"userName": userName,
		})
	}
}
