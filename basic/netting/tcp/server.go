package tcp

import (
	"HelloGo/basic/body"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func transfer(conn net.Conn) {
	defer func() {
		remoteAddr := conn.RemoteAddr().String()
		log.Print("discard remove add:", remoteAddr)
		conn.Close()
	}()

	for {
		var msg body.Message

		if err := json.NewDecoder(conn).Decode(&msg); err != nil && err != io.EOF {
			log.Printf("Decode from client err: %v", err)
			return
		}

		if msg.Uid != "" || msg.Val != "" {
			//conn.Write([]byte(msg.Val))
			ser, _ := json.Marshal(msg)
			conn.Write(append(ser, '\n'))
		}
	}
	//buf := make([]byte, 1024)
	//// Read the incoming connection into the buffer.
	//_, err := conn.Read(buf)
	//if err != nil && err != io.EOF {
	//	log.Printf("read from client err: %v", err)
	//	return
	//}
	//
	//json.Unmarshal(buf, msg)

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
		//go handle(conn)
		go transfer(conn)
	}
}
