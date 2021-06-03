package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

const (
	CON = 10
)

func main() {
	//runtime.GOMAXPROCS(1)
	var eg errgroup.Group
	//eg, errCtx := errgroup.WithContext(context.Background())
	for i := 0; i < CON; i++ {
		i := i
		eg.Go(func() error {
			time.Sleep(1 * time.Second)
			if i == 3 {

				return errors.New("Mock err!")
			}
			fmt.Println(i)
			return nil
		})
	}

	// wait for done or err occurs
	if err := eg.Wait(); err == nil {
		log.Print("all sounds good.")
	} else {
		log.Printf("errors occurs: %v", err)
	}
	return
}
