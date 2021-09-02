package middle

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

func CacheForReq(expire time.Duration, handle gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
