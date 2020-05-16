package main

import (
	"fmt"
	"runtime"
	"time"
)

//通知管道
var sig = make(chan bool)

func main()  {
	defer fmt.Println("Main routines exit!")
	//FeatureExit()
	ExitBySignal()
	fmt.Println("Start with goroutines num:", runtime.NumGoroutine())
	sig <- true
	fmt.Println("Before finished, goroutines num:", runtime.NumGoroutine())
}

func ExitBySignal()  {
	go ListenWithSignal()
	time.Sleep(time.Second)
}

//利用管道通知协程退出
func ListenWithSignal()  {
	count := 1
	for {
		select {
		//监听通知
		case <-sig:
			return
		default:
			//正常执行
			time.Sleep(100 * time.Millisecond)
			count++
		}
	}
}

// 让子协程自生自灭
func FeatureExit()  {
	//新起子协程
	go Spawn()
	time.Sleep(time.Second)
}

func Spawn()  {
	count := 1
	for {
		time.Sleep(100 * time.Millisecond)
		count++
	}
}
