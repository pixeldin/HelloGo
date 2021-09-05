package cache

import "time"

// 定义缓存接口
type CacheStore interface {
	// 获取指定key
	Get(key string, value interface{}) error
	// 设置指定key
	Set(key string, value interface{}, expire time.Duration) error
	// 删除指定key
	Delete(key string) error
	// 清除所有key
	Flush() error
}
