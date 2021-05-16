package main

import (
	"log"
	"sync"
	"time"
)

type signal struct {
	ready bool
}

func main() {
	mtx := new(sync.Mutex)
	cond := sync.NewCond(mtx)
	signal := new(signal)
	go ForWhileSignal(signal, "G-等待方", mtx, cond)
	time.Sleep(100)
	go Condition(signal, "G-唤醒方", mtx, cond)

	for {
		time.Sleep(1)
	}
}

func Condition(sig *signal, gid string, mutx *sync.Mutex, cd *sync.Cond) {
	mutx.Lock()
	defer func() {
		log.Print(gid + " 执行defer()\n")
		mutx.Unlock()
	}()

	log.Print(gid + " do something...\n")
	time.Sleep(3 * time.Second)
	sig.ready = true
	log.Print(gid + " 唤醒等待方...\n")
	cd.Signal()
}

func ForWhileSignal(sig *signal, gid string, mutx *sync.Mutex, cd *sync.Cond) {
	mutx.Lock()
	defer func() {
		log.Print(gid + " 执行defer()\n")
		mutx.Unlock()
	}()

	for !sig.ready {
		log.Print(gid + " 等待唤醒...\n")
		cd.Wait()
	}

	log.Print(gid + " 被叫醒! \n")
}
