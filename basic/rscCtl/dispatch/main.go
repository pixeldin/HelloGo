package main

import (
	"HelloGo/basic/rscCtl/dispatch/runner"
	"log"
	"os"
	"time"
)

//Runner调度器的Demo

//运行时间限制
const timeout = 3 * time.Second

func main() {
	log.Println("Start work.")

	r := runner.New(timeout)

	r.Add(TaskMaker(), TaskMaker(), TaskMaker())

	err := r.Start()
	if err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Stop cause system interrupted.")
			os.Exit(2)
		}
	}

	log.Println("Tasks finished.")
}

//返回函数任务
func TaskMaker() func(int) {
	return func(i int) {
		log.Printf("Processor - Task No.%d", i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}
