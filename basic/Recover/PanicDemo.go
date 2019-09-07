package main

import (
	"errors"
	"fmt"
)

func main() {
	//Try to recover fatal err
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recover panic: ", p)
		}
	}()

	fmt.Println("Call some wrong.")
	PanicCall()
}


func PanicCall()  {
	somethingWrong()
}

func somethingWrong()  {
	panic(errors.New("Fatal error!"))
}
