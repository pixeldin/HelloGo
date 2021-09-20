package cache

import (
	"HelloGo/common/model"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
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

func CacheForReq(store CacheStore, expire time.Duration, handle gin.HandlerFunc) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		/*
			拼装业务key
			从context获取请求体, 从store中根据key判断缓存是否命中
			- 命中, 返回
			- 不命中, 执行handler并写入缓存
		*/
		var (
			tryCache respCache
		)
		// todo... 拼接key
		url := ginCtx.Request.URL
		bodyStr := dumpForBody(ginCtx.Request)
		key := url.RequestURI() + ":" + bodyStr
		if err := store.Get(key, &tryCache); err != nil {
			// 未命中
			handle(ginCtx)
		} else {
			resp := model.Response{}
			if err := json.Unmarshal(tryCache.Data, &resp); err != nil {
				ginCtx.JSON(model.Unknown, model.NewErrorResponse(errors.New("cache unmarshal failed")))
				return
			}
			ginCtx.JSON(http.StatusOK, &resp)
		}
	}
}

// dumpForBody
// 从request体(ReadCloser)导出body str
func dumpForBody(request *http.Request) string {
	var bout bytes.Buffer
	if request.Body != nil {
		_, err := io.Copy(&bout, request.Body)
		if err != nil {

		}
		request.Body.Close()
	}
	// todo... fixme for correct key
	body := bout.String()
	return body
}
