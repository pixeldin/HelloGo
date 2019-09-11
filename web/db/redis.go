package db

import (
	"github.com/gomodule/redigo/redis"
	"github.com/obase/conf"
	"time"
)

const CONF_KEY  = "redis"

type RedisConf struct {
	address string
	pwd string
}

var rc RedisConf

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
	configs, ok := conf.GetSlice(CONF_KEY)
	if !ok || len(configs) == 0 {
		return
	}
	for _, config := range configs {
		addr, ok := conf.ElemString(config, "address")
		if ok {
			rc.address = addr
		}
		pwd, ok := conf.ElemString(config, "password")
		if ok {
			rc.pwd = pwd
		}

	}
	pool = NewPool(rc.address)
}

func GetRedisConn() (conn redis.Conn, err error){
	conn = pool.Get()
	_, err = conn.Do("AUTH", rc.pwd)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
