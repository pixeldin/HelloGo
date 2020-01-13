package worker

import (
	"fmt"
	. "github.com/sirupsen/logrus"
	"time"
)

func Produce(PNum int, msg chan []byte, signal <-chan int) {
	//周期生产
	PeriodJob(msg, signal, func() {
		cake := []byte(fmt.Sprintf("cake, @tag: PNum.%d", PNum))
		msg <- cake
		Infof("## Producer push a %s", cake)
	})
}

func PeriodJob(msg chan []byte, conn <-chan int, job func()) bool {
	for {
		//timer := time.NewTimer(1 * time.Second)
		tk := time.Tick(1 * time.Second)
		select {
		case <-conn:
			//timer.Stop()
			Info("Receive Produce stop signal.")
			return true
		//case <-timer.C: // 超时
		case <-tk:
			job()
			//timer.Stop()
		}
	}
}
