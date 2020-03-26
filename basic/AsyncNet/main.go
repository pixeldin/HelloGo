package main

import (
	"encoding"
	"log"
	"net"
)

func main() {
	l, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal(e)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go doSomeThing(conn)
	}
}

func doSomeThing(conn net.Conn)  {
	defer conn.Close()
	for {
		conn.SetReadDeadline()
	}
}
