package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	//每一层循环泄漏两个协程
	for i := 0; i < 4; i++ {
		//LeakSomeRoutine()
		FixLeakingByContex()
		//异步清理协程
		time.Sleep(100)
		fmt.Printf("#Goroutines in roop end: %d.\n", runtime.NumGoroutine())
	}
}

//泄漏协程demo
func LeakSomeRoutine() int {
	ch := make(chan int)
	//起3个协程抢着输入到ch
	go func() {
		ch <- 1
	}()

	go func() {
		ch <- 2
	}()

	go func() {
		ch <- 3
	}()
	//一有输入立刻返回
	return <-ch
}

func FixLeakingByContex() {

	//创建上下文用于管理子协程
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)
	go CancelByContext(ctx, ch)
	go CancelByContext(ctx, ch)
	go CancelByContext(ctx, ch)

	// 触发某个子协程退出
	ch <- 1
	//清理未结束协程
	defer cancel()
}

func CancelByContext(ctx context.Context, ch chan (int)) int {
	select {
	case <-ctx.Done():
		//fmt.Println("cancel by ctx.")
		return 0
	case n := <-ch :
		return n
	}
}
