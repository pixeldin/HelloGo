package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestSpawn(t *testing.T) {
	//打印已有协程数量
	fmt.Println("Start with goroutines num:", runtime.NumGoroutine())
	//新起子协程
	go Spawn()
	time.Sleep(time.Second)
	fmt.Println("After spawn goroutines num:", runtime.NumGoroutine())
	fmt.Println("Main routines exit!")
}
