package cache

import "time"

// storeI接口的实现体, 由外部创建实例
type InMemoryStore struct {
	CacheStore
}

func NewInMemoryStore(exp time.Duration) *InMemoryStore {
	// fixme... import memory-cache
	return &InMemoryStore{}
}
