package main

import (
	"fmt"
	"strconv"
)

type Pks struct {
	i  int
	gn string
}

func main() {
	n := 10
	c := make(chan Pks)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < n; j++ {
				pk := Pks{i, strconv.Itoa(j)}
				c <- pk
				fmt.Println("Goroutine ", i, " with task ", j)
			}
			done <- true
		}()
	}

	//Wait above to finish
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
		//close(done)
	}()

	for n := range c {
		fmt.Println("Pk No.", n.i, ",task:", n.gn)
	}

}
