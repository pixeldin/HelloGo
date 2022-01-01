package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	server http.Server
)

// 优雅停止demo
func main() {
	// 注册返回绑定了os.Interrupt的ctx
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server = http.Server{
		Addr: ":8080",
	}

	// 注册路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 10)
		fmt.Fprint(w, "Hello World")
		// diff
	})

	// 启动监听
	go server.ListenAndServe()

	// 触发interrupt信号
	<-ctx.Done()

	// 解绑上下文与信号量
	stop()
	log.Print("接收到SIGINT信号, 执行优雅停止, 等待收尾...")

	// 最后5秒回收连接
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	if err := server.Shutdown(timeoutCtx); err != nil {
		var TAG = "优雅停止错误: "
		log.Print(TAG, err)
		return
	}

	log.Print("程序关闭完成.")
}
