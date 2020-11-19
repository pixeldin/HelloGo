package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main4() {
	baton := make(chan int)

	waitGroup.Add(1)

	//Ready to run
	go Runner(baton)

	//Start to run
	baton <- 1
	waitGroup.Wait()

}

func Runner(baton chan int) {
	cicle := <-baton
	fmt.Printf("Runner %d run with Baton\n", cicle)
	nextRunner := cicle + 1

	if cicle < 4 {
		fmt.Println("Runner ", nextRunner, "ready to go...")
		go Runner(baton)
	}

	time.Sleep(100)

	if cicle == 4 {
		fmt.Println("Last cicle, done.")
		waitGroup.Done()
		return
	}

	fmt.Println("Runner ", cicle, " ok, to next cicle")
	baton <- nextRunner

}
