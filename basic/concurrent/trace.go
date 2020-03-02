package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func main() {

	//runtime.GOMAXPROCS(1)
	//go func() {
	//	for {
	//		//保持空转,如果没有其他函数调用,在单个处理器下该协程不会切换
	//		;
	//
	//		//函数调用, 触发协程切换(主协程)
	//		//Callsomething(total)
	//		//runtime.Gosched()
	//	}
	//}()

	//尝试切换协程
	//time.Sleep(100)
	//println("main done.")

	TraceDemo()
}

func TraceDemo() {
	/*
		trace 工具分析 go1.12
		go run trace.go 2> trace.out
		go tool trace trace.out
	*/
	trace.Start(os.Stderr)
	defer trace.Stop()

	var wg sync.WaitGroup
	var unsafeA = 0
	var safeA = 0
	var mut sync.Mutex

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func() {

			time.Sleep(7)
			unsafeA++

			mut.Lock()
			safeA++
			mut.Unlock()
			//http.Get(`https://httpstat.us/200?sleep=10000`)
			//fmt.Printf("Do something of i: %d\n", i)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(unsafeA)
	fmt.Println(safeA)
}
