package middle

import (
	"HelloGo/basic/http/kissgin/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"reflect"
)

func HeaderCheck(key string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取header的值
		kh := c.GetHeader(key)
		if kh == "" {
			// header缺失
			c.JSON(http.StatusOK, &model.Response{Code: model.Unknown, Msg: "lacking necessary header"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func ReqCheck(reqVal interface{}) func(ctx *gin.Context) {
	var reqType reflect.Type = nil
	// 运行时: 拿到校验体原始类型
	if reqVal != nil {
		value := reflect.Indirect(reflect.ValueOf(reqVal))
		reqType = value.Type()
	}

	return func(c *gin.Context) {
		tag := c.Request.RequestURI

		var req interface{} = nil
		if reqType != nil {
			req = reflect.New(reqType).Interface()
			// 类型校验
			if err := c.ShouldBindBodyWith(req, binding.JSON); err != nil {
				// 结构体绑定出错
				c.JSON(http.StatusOK, model.NewBindFailedResponse(tag))
				// 终止执行链
				c.Abort()
				return
			}
		}
		// 无需校验, 执行链往下
		c.Next()
	}
}
