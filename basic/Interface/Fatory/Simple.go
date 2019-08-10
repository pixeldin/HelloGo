package Fatory

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")

	readString, e := inputReader.ReadString('\n')
	if e != nil {
		fmt.Printf("An error occurred: %s\n", e)
		os.Exit(1)
	}

	name := readString[:len(readString)-1]
	fmt.Printf("Hello, %s. What can I do for you?\n", name)

	for {
		s, e := inputReader.ReadString('\n')
		if e != nil {
			fmt.Printf("An error occurred: %s\n", e)
			os.Exit(1)
		}
		s = s[:len(s) - 1]
		s = strings.ToLower(s)
		switch s {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("bye!")
			os.Exit(0)
		default:
			fmt.Println("Sorry, I can't not do that.")
		}
	}
}
