package controller

import (
	"github.com/gin-gonic/gin"
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
		//redirect to index
		//c.Redirect(http.StatusFound, "https://www.baidu.com")
		c.Redirect(http.StatusMovedPermanently, "manager/index")
	} else {
		c.JSON(200, gin.H{
			"msg":      "Auth failed",
			"userName": userName,
		})
	}
}
