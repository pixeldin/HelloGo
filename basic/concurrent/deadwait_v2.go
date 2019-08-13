package main

import (
	"fmt"
)

//WithSemaphore.go
type Pk struct {
	i int
	gn string
}

func main()  {

	c := make(chan Pk, 2)
	//wg.Add(2)
	done := make(chan bool)

	//G1
	go func() {
		//wg.Add(1)
		for i := 0; i < 10; i++ {
			pk := Pk{i, "G1"}
			c <- pk
		}
		fmt.Println("G1 done")
		//wg.Done()
		done <- true
	}()

	//G2
	go func() {
		//wg.Add(1)
		for i := 0; i < 10; i++ {
			pk := Pk{i, "G2"}
			c <- pk
		}
		fmt.Println("G2 done")
		//wg.Done()
		done <- true
	}()

	//go function() {
	//	wg.Wait()
	//	close(c)
	//}()

	//go function() {
	//	<-done
	//	<-done
	//	close(c)
	//}()

	//block here, unless below cover, but next won't cover cause above block G1, G2
	<-done
	<-done
	close(c)


	for n := range c {
		fmt.Println("Pop chan with : ", n.gn, "i: ", n.i)
	}


}
