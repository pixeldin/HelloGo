package main

import (
	"bytes"
	"os"
	"sync"
)

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg sync.WaitGroup
}

func main() {
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	b.WriteTo(os.Stdout)

}