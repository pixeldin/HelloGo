package cbk

import (
	"HelloGo/basic/cbk/util"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

// 请求API快照
type apiSnapShop struct {
	isPaused   bool  // api是否熔断
	errCount   int64 // api在周期内失败次数
	totalCount int64 // api在周期内总次数

	accessLast int64 // api最近一次访问时间
	roundLast  int64 // 熔断器周期时间
}

// 熔断器实现体
type CircuitBreakerImp struct {
	lock            sync.RWMutex
	apiMap          map[string]*apiSnapShop // api集合
	minCheck        int64                   // 接口熔断开启下限次数
	cbkErrRate      float64                 // 接口熔断开启比值
	recoverInterval time.Duration           // 熔断恢复区间
	roundInterval   time.Duration           // 计数重置区间
}

// accessed 记录访问
func (c *CircuitBreakerImp) accessed(api *apiSnapShop) {
	/*
		判断是否大于周期时间
		- 是: 重置计数
		- 否: 更新计数
	*/
	now := time.Now().UnixNano()
	if util.Abs64(now-api.roundLast) > int64(c.roundInterval) {
		log.Warnf("# Cbk reset for all keys in new round!")
		api.errCount = 0
		api.totalCount = 0
		api.roundLast = now
	}
	api.totalCount++
	api.accessLast = now
}

// CanAccess 判断api是否可访问
func (c *CircuitBreakerImp) CanAccess(key string) bool {
	/*
		判断当前api的isPaused状态
		- 未熔断, 返回true
		- 已熔断, 当前时间与恢复期比较
			- 大于恢复期, 返回true
			- 小于恢复期, 返回false
	*/
	c.lock.RLock()
	defer c.lock.RUnlock()
	// 从api全局map查找
	if api, ok := c.apiMap[key]; ok {
		if api.isPaused {
			// 判断是否进入恢复期
			latency := util.Abs64(time.Now().UnixNano() - api.accessLast)
			if latency < int64(c.recoverInterval) {
				// 在恢复期之内, 熔断
				return false
			}
		}
	}
	return true
}

/*
	Failed 记录失败访问
	api列表查找,
		- 已有:
			- 记录访问/错误次数
			- 是否失败占比到达阈值? 是, 则标记置为熔断
		- 未找到:
			更新至api列表: 记录访问/错误次数
*/
func (c *CircuitBreakerImp) Failed(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if api, ok := c.apiMap[key]; ok {
		c.accessed(api)
		api.errCount++

		errRate := float64(api.errCount) / float64(api.totalCount)
		// 请求数量达到阈值 && 错误率高于熔断界限
		if api.totalCount > c.minCheck && errRate > c.cbkErrRate {
			log.Warnf("Cbk start for key: %v", key)
			api.isPaused = true
		}
	} else {
		api := &apiSnapShop{}
		c.accessed(api)
		api.errCount++
		// 写入全局map
		c.apiMap[key] = api
	}
}

/*
	Succeed 记录成功
	只更新api列表已有的,
	记录访问, 并判断是否熔断:
	- 是, 取消熔断状态
*/
func (c *CircuitBreakerImp) Succeed(key string) {
	c.lock.Lock()
	c.lock.Unlock()

	if api, ok := c.apiMap[key]; ok {
		c.accessed(api)
		if api.isPaused {
			api.isPaused = false
		}
	}
}

// IsBreak 判断api熔断状态
func (c *CircuitBreakerImp) IsBreak(key string) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if api, ok := c.apiMap[key]; ok {
		return api.isPaused
	}
	return false
}

func (c *CircuitBreakerImp) Status() interface{} {
	panic("implement me")
}
