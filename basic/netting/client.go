package netting

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func RegistAndSend() {
	conn, e := net.Dial("tcp", "127.0.0.1:3000")
	if e != nil {
		fmt.Println("Err while dialing:", e.Error())
		return
	}
	defer conn.Close()

	//Wait for msg from client
	inputReader := bufio.NewReader(os.Stdin)
	for {
		//Send something to server
		str, _ := inputReader.ReadString('\n')
		msg := strings.Trim(str, "\n")
		if msg == "quit" {
			return
		}
		_, err := conn.Write([]byte(str))
		if err != nil {
			fmt.Println("Send data fail", err)
			return
		}

		//Read data from server
		buff := make([]byte, 512)
		n, err := conn.Read(buff)
		fmt.Println("Data from server: ", string(buff[:n]))
	}
}

//seems not working in unit test
func InputExa() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}
	}
}
