package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var mtx = new(sync.Mutex)
var c *sync.Cond = sync.NewCond(mtx)

func TestPV(t *testing.T) {
	go SignalCron()
	// just wait
	for {
		mtx.Lock()
		fmt.Println("Ready to wait...")
		c.Wait()
		fmt.Println("Weak up")
		time.Sleep(1 * time.Second)
		mtx.Unlock()
	}
}

func SignalCron() {
	for {
		mtx.Lock()
		time.Sleep(3 * time.Second)
		c.Signal()
		mtx.Unlock()
	}
}
