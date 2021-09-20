package main

import (
	"HelloGo/basic/http/kissgin/bizmod"
	"HelloGo/basic/http/kissgin/middle"
	"HelloGo/basic/http/kissgin/middle/cache"
	comdel "HelloGo/common/model"
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
	H_KEY    = "h-key"
	CacheExp = time.Second * 30
)

func main() {
	// 30s缓存中间件
	store := cache.NewInMemoryStore(CacheExp)
	r := gin.Default()
	//r.Use(middle.HeaderCheck(), middle.ReqCheck())

	r.POST("/hello", helloFunc)
	// 校验header
	r.POST("/hello-with-header", middle.HeaderCheck(H_KEY), helloFunc)
	// 检测请求体
	r.POST("/hello-with-req", middle.ReqCheck(bizmod.PingReq{}), helloFunc)

	// todo... 接口缓存cache, 对相同uri,相同参数生效
	r.POST("/hello-with-cache", cache.CacheForReq(store, CacheExp, helloFunc))

	e := r.Run()
	fmt.Printf("Server stop with err: %v\n", e)
}

func helloFunc(c *gin.Context) {
	const TAG = "PingPong"
	c.JSON(comdel.Success, comdel.SimpleResponse(200, TAG))
	return
}
