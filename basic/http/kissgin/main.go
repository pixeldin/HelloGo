package main

import (
	"HelloGo/basic/http/kissgin/middle"
	"HelloGo/basic/http/kissgin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func init() {
	// 设置gin启动默认端口
	if err := os.Setenv("PORT", "8099"); err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()
	//r.Use(middle.HeaderCheck(), middle.ReqCheck())

	r.POST("/hello", helloFunc)
	// 中间件校验header
	r.POST("/hello-with-header", middle.HeaderCheck(), helloFunc)
	// 中间件检测请求体
	r.POST("/hello-with-req", middle.ReqCheck(model.PingReq{}), helloFunc)

	e := r.Run()
	fmt.Printf("Server stop with err: %v\n", e)
}

func helloFunc(c *gin.Context) {
	const TAG = "PingPong"
	c.JSON(http.StatusOK, model.SimpleResponse(200, TAG))
	return
}
