package netting

import (
	"HelloGo/basic/netting/http"
	"HelloGo/basic/netting/tcp"
	"testing"
)

func TestListenAndServer(t *testing.T) {
	tcp.ListenAndServer()
	//time.Sleep(3000)
	//RegistAndSend()
}

func TestRegistAndSend(t *testing.T) {
	tcp.RegistAndSend()
}

func TestInputExa(t *testing.T) {
	tcp.InputExa()
}

func TestStartHelloWithHttp(t *testing.T)  {
	http.StartHelloWithHttp()
}
