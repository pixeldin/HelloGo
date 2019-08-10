package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	newReader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input your name:")
	readString, e := newReader.ReadString('\n')
	if e != nil {
		fmt.Print("Err: %s\n", e)
	} else {
		input := readString[:len(readString)-1 ]
		fmt.Printf("Hello, %s!", input)
	}

}
