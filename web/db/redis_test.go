package db

import (
	"HelloGo/web/common"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestGetRedisConn(t *testing.T) {
	conn := GetRedisConn()
	defer conn.Close()

	conn.Do("AUTH", "pixelpig")
	value, e := redis.String(conn.Do("get", "pixel"))
	common.Logging(common.ErrCheck("Redis execute err", e))
	fmt.Println((value))
}
