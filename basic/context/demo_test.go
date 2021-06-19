package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestSpawn(t *testing.T) {
	//打印已有协程数量
	fmt.Println("Start with goroutines num:", runtime.NumGoroutine())
	//新起子协程
	go Spawn()
	time.Sleep(time.Second)
	fmt.Println("After spawn goroutines num:", runtime.NumGoroutine())
	fmt.Println("Main routines exit!")
}

func hardWork(job interface{}) error {
	time.Sleep(3 * time.Second)
	return nil
}

func requestWork(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	done := make(chan error)
	go func() {
		done <- hardWork(job)
	}()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func TestRw(t *testing.T) {
	const total = 1 //000
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()
	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			requestWork(context.Background(), "any")
		}()
	}
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	time.Sleep(time.Second * 2)
	fmt.Println("number of goroutines:", runtime.NumGoroutine())
}
