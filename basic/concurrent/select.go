package main

import (
	"fmt"
	"time"
)

func main() {
	go leakRoutine()
	time.Sleep(3000)
	fmt.Println("Until Main exit, routine before will leak in memory")
}

//该协程会阻塞, 造成泄露, 如果加进主协程,会导致死锁
func leakRoutine() {
	//没有初始化的string通道
	var strings <-chan string
	for s := range strings {
		fmt.Println(s)
	}
}
