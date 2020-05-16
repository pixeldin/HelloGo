package selfDef

import "context"

type PixelCtx struct {
	context.Context
	context.CancelFunc

	//全局键值对, 可以管理全局配置
	conf map[string]interface{}
	//虚拟资源
	mongoDB string
	mysqlDB string
	redisDB string
}
