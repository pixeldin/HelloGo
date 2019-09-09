package db

import (
	"HelloGo/web/common"
	"github.com/gomodule/redigo/redis"
	"fmt"
	"testing"
)

func TestGetRedisConn(t *testing.T) {
	conn := GetRedisConn()
	defer conn.Close()

	conn.Do("AUTH", "sgpsvr")
	value, e := redis.String(conn.Do("get", "pixel"))
	if e != nil {
		common.Logging(common.ErrCheck("Redis execute err", e))
		return
	}
	fmt.Println((value))
}
