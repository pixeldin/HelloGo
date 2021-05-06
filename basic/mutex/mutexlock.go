package main

import (
	"fmt"
	"sync"
)

type service struct {
	inProgress         map[string]bool
	awaitingCompletion map[string][]chan string
	lock               sync.RWMutex
}

func main_l() {
	s := "hello world, 宇宙"
	trs := []rune(s)
	fmt.Println(trs)
	by := []byte(s)
	fmt.Println(by)

	//mp := make(map[string]int, 10)
	mp := map[string]int{}
	mp["0"] = 0
	mp["1"] = 1
	mp["2"] = 2
	mp["3"] = 3
	for k, v := range mp {
		fmt.Print(k, "r:", v)
	}
	fmt.Println()
	for k, _ := range mp {
		fmt.Print(k, "v:", mp[k])
	}
}
