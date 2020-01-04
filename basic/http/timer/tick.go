package timer

import (
	"fmt"
	"time"
)

func PeriodJob(conn <-chan string) bool {
	for {
		timer := time.NewTimer(1 * time.Second)
		select {
		case <-conn:
			timer.Stop()
			println("Time ticking done.")
			return true
		case <-timer.C: // 超时
			println("Do something after period round.")
			timer.Stop()
		}
	}
}

func Tick()  {
	tick := time.Tick(5 * time.Second)
	for now := range tick {
		fmt.Println(now)
	}
}
