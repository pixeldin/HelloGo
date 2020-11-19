package main

import (
	"fmt"
	"sync"
)

func Main()  {
	type Pk struct {
		i int
		gn string
	}
	c := make(chan Pk)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			pk1 := Pk{i, "G1"}
			c <- pk1
		}
		fmt.Println("G1 done")
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			pk2 := Pk{i, "G2"}
			c <- pk2
		}
		fmt.Println("G2 done")
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()


	for n := range c {
		fmt.Println("Pop chan with : ", n.gn, "i: ", n.i)
	}

}
