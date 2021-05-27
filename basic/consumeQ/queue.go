package consumeQ

import (
	"fmt"
	"log"
	"time"
)

var QUEUE_SIZE = 10

const (
	GLB_KEY = "RDS_CHAT_MSG_BACK_QUEUE"
)

func MqBackupToMongo() {

	msgIdChannel := make(chan string, QUEUE_SIZE)

	// 生产端
	//safego.Go(func() {
	go func() {
		//应对redis重启的情况
		for {
			//ctx := context.Background()
			count := 0
			func() {
				// protect inside
				fmt.Println("start produce...")
				defer func() {
					// 恢复到上一个func
					if p := recover(); p != nil {
						log.Print("Recover panic: ", p)
					}
				}()

				for {
					msgIdChannel <- "hello"
					time.Sleep(100)
					count++
					if count == 15 {
						panic("haha")
						//break
						time.Sleep(10 * time.Second)
						count = 0
					}
				}
			}()

			fmt.Println("just ticking...")
			// 保护性休眠
			time.Sleep(time.Second)
		}
	}()

	// 消费端
	//safego.Go(func() {
	go func() {
		msgIds := make([]string, 0)
		timer := time.NewTimer(time.Second * 3)
		defer timer.Stop()

		for {
			select {
			case msgId := <-msgIdChannel:
				msgIds = append(msgIds, msgId)
				if len(msgIds) >= QUEUE_SIZE {
					// process msgIds with batch
					fmt.Println("Process batch msg with size: ", len(msgIds))
					msgIds = make([]string, 0)
				}
				// 重置计时器
				timer.Reset(time.Second * 3)
			case <-timer.C:
				// 首次或者在每次Reset之后的3秒执行一次收尾工作
				if len(msgIds) > 0 {
					// process msgIds with less
					fmt.Println("Timer: Process batch msg with size: ", len(msgIds))
					msgIds = make([]string, 0)
				}
			}
		}
	}()
}
