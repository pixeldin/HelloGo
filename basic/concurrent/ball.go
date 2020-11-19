package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


var wg sync.WaitGroup

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func main2()  {
	fmt.Println("BaseBall start.")

	court := make(chan int)
	wg.Add(2)
	go player("A", court)
	go player("B", court)

	court <- 1
	wg.Wait()

}

func player(name string, court chan int)  {
	defer wg.Done()

	for {
		ball, ok := <- court
		if !ok {
			fmt.Println("Winner: ", name)
			return
		}

		n := rand.Intn(100)
		if n % 7 == 0 {
			fmt.Println("Player ", name, "lost key point...")
			//close(court)
			//lead the goroutine
			panic(court)
			return
		}
		fmt.Println("Player ", name, "Hit ", ball)
		ball++
		court <- ball
	}
}