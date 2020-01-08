package main

import (
	"HelloGo/transport/model"
	"HelloGo/transport/worker"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrus.Info("Start demo..., time: %v", time.Now())
	var costSize, prodSize = model.CONSUME_SIZE, model.PRODUCE_SIZE
	msgQueue := make(chan []byte, prodSize)
	ring := make(chan int)
	//Consumer job
	for i := 0; i < costSize; i++ {
		go worker.Consume(i, msgQueue)
	}

	//Producer job
	for i := 0; i < costSize; i++ {
		go worker.Produce(i, msgQueue, ring)
	}

	time.Sleep(5 * time.Second)
	ring <- 1
	logrus.Info("Main job done!")
}
