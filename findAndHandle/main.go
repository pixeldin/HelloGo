package main

import (
	"HelloGo/findAndHandle/model"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {

	OutputStatic()

	defer func() {
		log.Info("Process done.")
	}()

	go process()
}

func process()  {
	for  {
		if err := findAndHandle(); err != nil {
			log.Errorf("Process err %v", err)
		}
		time.Sleep(model.PROCESS_TICK_TIME)
	}
}

func findAndHandle() error {
	//轮询获取事件通知

	//获取信号加锁

	//处理任务

	//标识事件状态

	return nil
}

func OutputStatic()  {
	loca, e := os.Executable()
	if e != nil {
		log.Infof("e: %v", e)
	}
	log.Infof("Ready to process from %v...", loca)

	dir, e := os.Getwd()
	if e != nil {
		log.Infof("e: %v", e)
	}
	log.Infof("File wd %v...", dir)
}