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
	go ForWhileSignal(signal, "G-FW", mtx, cond)
	time.Sleep(100)
	go Condition(signal, "G-HD", mtx, cond)

	for {
		time.Sleep(1)
	}
}

func Condition(sig *signal, gid string, mutx *sync.Mutex, cd *sync.Cond) {
	mutx.Lock()
	defer func() {
		log.Print(gid + " defer...\n")
		//cd.Signal()
		mutx.Unlock()
	}()

	log.Print(gid + " do something...\n")
	time.Sleep(3 * time.Second)
	sig.ready = true
	log.Print(gid + " call another...\n")
	cd.Signal()

}

func ForWhileSignal(sig *signal, gid string, mutx *sync.Mutex, cd *sync.Cond) {
	mutx.Lock()
	defer func() {
		log.Print(gid + " defer...\n")
		mutx.Unlock()
	}()

	for !sig.ready {
		log.Print(gid + " wait till signal be awaken...\n")
		cd.Wait()
	}

	log.Print(gid + " be awaken! \n")
}
