package tcp

import (
	"HelloGo/basic/body"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"time"
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

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var buffer []byte = []byte("You are welcome. I'm server.")

	for {

		time.Sleep(1 * time.Second)
		n, err := conn.Write(buffer)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
		log.Println("send:", n)

		select {}
	}

	log.Println("connetion end")

}

const TAG = "server: hello, "

func transfer(conn net.Conn) {
	defer func() {
		remoteAddr := conn.RemoteAddr().String()
		log.Print("discard remove add:", remoteAddr)
		conn.Close()
	}()

	// 设置10秒关闭连接
	//conn.SetDeadline(time.Now().Add(10 * time.Second))

	for {
		var msg body.Message

		if err := json.NewDecoder(conn).Decode(&msg); err != nil && err != io.EOF {
			log.Printf("Decode from client err: %v", err)
			// todo... 仿照redis协议写入err前缀符号`-`，通知client错误处理
			return
		}

		if msg.Uid != "" || msg.Val != "" {
			//conn.Write([]byte(msg.Val))
			var rsp body.Resp
			rsp.Uid = msg.Uid
			rsp.Val = TAG + msg.Val
			rsp.Ts = time.Now().String()
			ser, _ := json.Marshal(msg)

			time.Sleep(3 * time.Second)
			conn.Write(append(ser, '\n'))
		}
	}
}

func ListenAndServer() {
	log.Print("Start server...")
	listen, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatal("Listen failed. msg: ", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept failed, err: %v", err)
			continue
		}
		go transfer(conn)
		//go handleConnection(conn)
	}
}
