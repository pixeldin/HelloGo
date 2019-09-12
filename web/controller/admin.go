package controller

import (
	"HelloGo/web/common"
	"HelloGo/web/constant"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RegisterAdminRouter(group *gin.RouterGroup) {
	a := group.Group("/manager")
	{
		a.GET("/index", Index)
		a.GET("/mockdata", MatchData)
		a.GET("/graph", GraphList)
		a.GET("/summary", SummaryTable)
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
		logrus.Debug("Nil user from session, redirect to login")
		c.HTML(http.StatusFound, "auth/login", h)
	} else {
		logrus.Debug("=====================Redirect with old session=====================")
		//jump to main page
		h["user"] = user.(string)
		c.HTML(http.StatusFound, "manager/index", h)
	}
}

type matchSum struct {
	MatchType int `json:"matchType"`
	ServerType string `json:"serverType"`
	SrcLog int  `json:"srcLog"`
} 

func MatchData(context *gin.Context) {
	//TODO: Session auth
	msa := []matchSum{}
	ms := matchSum{3, "c", 1000}
	ms2 := matchSum{2, "c", 10000}
	msa = append(msa, ms)
	msa = append(msa, ms2)
	context.JSON(200, gin.H{
		"code":0,
		"msg":"mock something data.",
		"count":10,
		//"data": "[{'matchType':3, 'serverType':'c'},{'SrcLog':1000}]",
		"data": msa,
		//"matchType": 3,
		//"serverType": "c",
		//"SrcLog": 100000,
		//"pvpResult": 100,
		//"aggData": 10,
		//"ckh": 10,
		//"date": "2019/09/11",
	})
}

func SummaryTable(ctx *gin.Context)  {
	h := gin.H{}
	ctx.HTML(http.StatusFound, "manager/summary", h)
}

func GraphList(ctx *gin.Context)  {
	//TODO: Session auth
	h := gin.H{}
	ctx.HTML(http.StatusFound, "manager/graph", h)
}
