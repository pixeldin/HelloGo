package main

import (
	"HelloGo/basic/http/timer"
	"fmt"
	"time"
)

func main() {

	defer fmt.Println("## Main finished.")

	chs := make(chan string)
	go timer.PeriodJob(chs)
	time.Sleep(5 * time.Second)
	chs <- "ok"

	//timer.Tick()

}

