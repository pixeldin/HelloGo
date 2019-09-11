package db

import (
	"HelloGo/web/common"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetRedisConn(t *testing.T) {
	conn, e := GetRedisConn()
	if conn != nil {
		defer conn.Close()
	}
	if e != nil {
		logrus.Error("Get redis conn failed, err: ", e)
		return
	}

	value, e := redis.String(conn.Do("get", "pixel"))
	if !common.LoggingErr(common.ErrCheck("Redis execute err", e)) {
		common.Infof("Get pixel value:", value)
	}
}
