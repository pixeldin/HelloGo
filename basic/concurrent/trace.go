package main

import (
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {

	//var total = new(int)

	runtime.GOMAXPROCS(1)
	go func() {
		for {
			//保持空转,如果没有其他函数调用,在单个处理器下该协程不会切换
			;

			//函数调用, 触发协程切换(主协程)
			//Callsomething(total)
			//runtime.Gosched()
		}
	}()

	//尝试切换协程
	time.Sleep(100)
	println("main done.")
}

func Callsomething(total *int) {
	println(*total)
}

func TraceDemo() {
	/*
		trace 工具分析 go1.12
		go run trace.go 2> trace.out
		go tool trace trace.out
	*/
	//trace.Start(os.Stderr)
	//defer trace.Stop()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			http.Get(`https://httpstat.us/200?sleep=10000`)
			//println("Do something")
			//time.Sleep(100)
			wg.Done()
		}()
	}

	wg.Wait()
}
