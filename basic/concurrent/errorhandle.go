package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	//checkStatus := func(done chan interface{}, urls ...string) <-chan Result { //2
	//	output := make(chan Result)
	//	go func() {
	//		defer close(output)
	//		for _, url := range urls {
	//			var result Result
	//			resp, err := http.Get(url)
	//			result = Result{Error: err, Response: resp} //3
	//			select {
	//			case <-done:
	//				return
	//			case output <- result: //4
	//			}
	//		}
	//	}()
	//	return output
	//}
	//done := make(chan interface{})
	//defer close(done)
	//
	//urls := []string{"https://www.baidu.com", "https://badhost"}
	//for status := range checkStatus(done, urls...) {
	//	if status.Error != nil {
	//		fmt.Println(status.Error)
	//		continue
	//	}
	//	fmt.Printf("Response : %v\n", status.Response.Status)
	//}

	msg := make(chan int)
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	fmt.Println("After 3 seconds, pass into msg.")
	//	msg <- 1
	//}()

	msg <- 1

	a := <-msg

	fmt.Printf("Got a %v, Main done!\n", a)
}
