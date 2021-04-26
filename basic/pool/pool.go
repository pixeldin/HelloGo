package pool

import (
	"container/list"
	"sync"
	"time"
)

type Option struct {
	addr        string
	size        int
	readTimeout time.Duration
	dialTimeout time.Duration
	keepAlive   time.Duration
}

type Pool struct {
	*Option
	idle    *list.List  // 空闲队列
	actives int         // 总连接数
	mutx    *sync.Mutex // 同步锁
	cond    *sync.Cond  // 用于阻塞/唤醒
}
