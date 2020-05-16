package netting

import (
	"HelloGo/basic/netting/tcp"
	"context"
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

func TestContext(t *testing.T) {
	// 获取顶级上下文
	ctx := context.Background()
	// 在上下文写入值, 注意需要返回新的value上下文
	valueCtx := context.WithValue(ctx, "hello", "pixel")
	value := valueCtx.Value("hello")
	if value != nil {
		fmt.Printf("Value type: %T/%v, value: %v.\n", value,
			reflect.TypeOf(value), value)
	}
}

//Test param from context
func TestRouterParamsFromContext(t *testing.T) {
	routed := false

	wantParams := httprouter.Params{httprouter.Param{"name", "gopher"}}
	handlerFunc := func(_ http.ResponseWriter, req *http.Request) {
		// get params from request context
		params := httprouter.ParamsFromContext(req.Context())

		if !reflect.DeepEqual(params, wantParams) {
			t.Fatalf("Wrong parameter values: want %v, got %v", wantParams, params)
		}

		routed = true
	}

	var nilParams httprouter.Params
	handlerFuncNil := func(_ http.ResponseWriter, req *http.Request) {
		// get params from request context
		params := httprouter.ParamsFromContext(req.Context())

		if !reflect.DeepEqual(params, nilParams) {
			t.Fatalf("Wrong parameter values: want %v, got %v", nilParams, params)
		}

		routed = true
	}
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/user", handlerFuncNil)
	router.HandlerFunc(http.MethodGet, "/user/:name", handlerFunc)

	w := new(mockResponseWriter)
	r, _ := http.NewRequest(http.MethodGet, "/user/gopher", nil)
	router.ServeHTTP(w, r)
	if !routed {
		t.Fatal("Routing failed!")
	}

	routed = false
	r, _ = http.NewRequest(http.MethodGet, "/user", nil)
	router.ServeHTTP(w, r)
	if !routed {
		t.Fatal("Routing failed!")
	}
}

func TestHttpRouter(t *testing.T) {
	router := httprouter.New()

	routed := false

	router.Handle("GET", "/user/ab/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//do nothing, just add path+handler
	})

	router.Handle("GET", "/user/abc/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//do nothing, just add path+handler
	})

	router.Handle("GET", "/user/a/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//do nothing, just add path+handler
	})

	router.Handle(http.MethodGet, "/user/abc/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
