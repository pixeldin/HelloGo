package main

import (
	"HelloGo/basic/http/kissgin/middle"
	"HelloGo/basic/http/kissgin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func init() {
	// 设置gin启动默认端口
	if err := os.Setenv("PORT", "8099"); err != nil {
		panic(err)
	}
}

var (
	H_KEY = "h-key"
)

func main() {
	r := gin.Default()
	//r.Use(middle.HeaderCheck(), middle.ReqCheck())

	r.POST("/hello", helloFunc)
	// 校验header
	r.POST("/hello-with-header", middle.HeaderCheck(H_KEY), helloFunc)
	// 检测请求体
	r.POST("/hello-with-req", middle.ReqCheck(model.PingReq{}), helloFunc)

	// todo... 接口缓存cache, 对相同uri,相同参数生效
	r.POST("/hello-with-cache", middle.CacheForReq(5*time.Minute, helloFunc))

	e := r.Run()
	fmt.Printf("Server stop with err: %v\n", e)
}

func helloFunc(c *gin.Context) {
	const TAG = "PingPong"
	c.JSON(model.Success, model.SimpleResponse(200, TAG))
	return
}
