package main

import (
	"fmt"
	"sync"
)

func main()  {
	type Pk struct {
		i int
		gn string
	}
	c := make(chan Pk, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			pk := Pk{i, "G1"}
			c <- pk
		}
		fmt.Println("G1 done")
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			pk := Pk{i, "G2"}
			c <- pk
		}
		fmt.Println("G2 done")
		wg.Done()
	}()

	//go func() {
	//	wg.Wait()
	//	close(c)
	//}()


	for n := range c {
		fmt.Println("Pop chan with : ", n.gn, "i: ", n.i)
	}

	defer close(c)
	wg.Wait()
}
