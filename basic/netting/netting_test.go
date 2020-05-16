package netting

import (
	"HelloGo/basic/netting/tcp"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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

func TestStartHelloWithHttp(t *testing.T) {
	//fmt.Println(path.Base(ul))

	ul := `https://localhost:8080/pixeldin/123`
	parse, e := url.Parse(ul)
	if e != nil {
		log.Fatalf("%v", e)
	}
	//fmt.Println(parse.Path)	// "/name/123"
	name := GetParamFromUrl(parse.Path, 1)
	id := GetParamFromUrl(parse.Path, 2)
	fmt.Println("name: " + name + ", id: " + id)
}

//指定返回相对url的值
func GetParamFromUrl(base string, index int) (ps string) {
	kv := strings.Split(base, "/")
	assert(index < len(kv), errors.New("index out of range."))
	return kv[index]
}

func assert(ok bool, err error) {
	if !ok {
		panic(err)
	}
}

func TestHttpRouter(t *testing.T) {
	router := httprouter.New()

	routed := false
	router.Handle("GET", "/user/ikbc", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//do nothing
	})
	router.Handle("GET", "/user/ikbcd", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//do nothing
	})
	router.Handle("GET", "/user/abc/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//do nothing
		ps.ByName("id")
	})

	router.Handle(http.MethodGet, "/pixel/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		routed = true
		want := httprouter.Params{httprouter.Param{"name", "gopher"}}
		if !reflect.DeepEqual(ps, want) {
			t.Fatalf("wrong wildcard values: want %v, got %v", want, ps)
		}
	})

	w := new(mockResponseWriter)

	req, _ := http.NewRequest(http.MethodGet, "/user/gopher", nil)
	router.ServeHTTP(w, req)

	if !routed {
		t.Fatal("routing failed")
	}
}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteHeader(s int) {

}