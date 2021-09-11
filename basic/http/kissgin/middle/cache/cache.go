package cache

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 需要缓存的结构
type respCache struct {
	Status int
	Header http.Header
	Data   []byte
}

// 用于覆盖gin.ResponseWrite在http返回操作处做缓存
type cacheWrite struct {
	gin.ResponseWriter
	status int           // 来自WriteHeader()的状态码
	store  CacheStore    // 传入缓存实例, 使用运行时缓存,如:go-cache 或者 外部缓存如:redis等
	key    string        // 缓存的键值, 可结合业务进行拼接, 如请求uri+body格式化的string
	expire time.Duration // respCache缓存过期时间
}

var _ gin.ResponseWriter = &cacheWrite{}

func CacheForReq(expire time.Duration, handle gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 从context获取请求体, 从store中判断缓存是否命中
		// - 命中, 返回
		// - 不命中, 执行handler并写入缓存
		handle(context)
	}
}
