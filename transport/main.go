package main

import (
	"HelloGo/transport/model"
	"HelloGo/transport/worker"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrus.Info("Start demo..., time: %v", time.Now())
	defer logrus.Info("Main job done!")
	//生产者,消费者并发度
	var costSize, prodSize = model.CONSUME_SIZE, model.PRODUCE_SIZE
	//消息管道
	msgQueue := make(chan []byte, model.QUEUE_SIZE)
	ring := make(chan int)
	//Consumer job
	for i := 0; i < costSize; i++ {
		go worker.Consume(i, msgQueue)
	}

	//Producer job
	for i := 0; i < prodSize; i++ {
		go worker.Produce(i, msgQueue, ring)
	}

	/*
		Try to idle
		注意: 如果协程处理时间大于主协程, 则有可能任务处理中断
	 */
	time.Sleep(3 * time.Second)
	close(ring)
}
