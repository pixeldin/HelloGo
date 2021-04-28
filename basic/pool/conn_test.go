package pool

import (
	"HelloGo/basic/body"
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestSome(t *testing.T) {
	var sm = new(sync.Map)
	sm.Store("hello", "world")
	if sm != nil {
		sm.Range(func(key, value interface{}) bool {
			fmt.Println(key)
			fmt.Println(value)
			return true
		})
	}
	//var val string
	val, ok := sm.Load("hello")
	if !ok {
		log.Fatal("not found key!")
	}
	log.Print(val.(string))
}

func TestSendMsg(t *testing.T) {
	opt := &Option{
		addr:        "0.0.0.0:3000",
		size:        3,
		readTimeout: 3 * time.Second,
		dialTimeout: 3 * time.Second,
		keepAlive:   30 * time.Second,
	}
	c, err := NewConn(opt)
	if err != nil {
		t.Fatal(err)
	}
	msg := &body.Message{Uid: "pixel-1", Val: "pixelpig!"}
	rec, err := c.Send(context.Background(), msg)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("rec1: %+v", <-rec)
	}

	msg.Val = "another pig!"
	rec2, err := c.Send(context.Background(), msg)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("rec2: %+v", <-rec2)
	}

}
