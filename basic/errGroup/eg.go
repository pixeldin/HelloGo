package main

import (
	"context"
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
	var eg *errgroup.Group
	eg, errCtx := errgroup.WithContext(context.Background())
	for i := 0; i < CON; i++ {
		i := i
		eg.Go(func() error {
			if i == 3 {
				//time.Sleep(2 * time.Second)
				return errors.New(fmt.Sprintf("Mock err: %d", i))
			}
			select {
			// 让Done分支判断在Mock err出现后命中, 等待1s
			case <-time.After(time.Duration(1) * time.Second):
			case <-errCtx.Done():
				log.Printf("meet err in job: %d", i)
				return errCtx.Err()
			}
			fmt.Println(i)
			return nil
		})
	}

	// wait for done or err occurs
	if err := eg.Wait(); err == nil {
		log.Print("all looks good.")
	} else {
		log.Printf("errors occurs: %v", err)
	}
	return
}
