package db

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func NewPool(addr string) *redis.Pool  {
	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 300 * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", addr)
		},
	}
}

func init()  {
	pool = NewPool("10.11.165.44:6379")
}

func GetRedisConn() redis.Conn {
	return pool.Get()
}
