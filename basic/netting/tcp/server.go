package tcp

import (
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 100)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		//read data from conn
		fmt.Println("Get data from conn, size: %d, msg: %s", n, string(buf[0:n]))
		msg := []byte("hello net\n.")
		conn.Write(msg)
	}
}
func ListenAndServer() {
	fmt.Println("Start server...")
	listen, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		fmt.Println("Listen failed. msg: ", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed")
			continue
		}
		go handle(conn)
	}
}
