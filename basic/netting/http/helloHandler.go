package http

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const form = `
	<html>
		<body>
			<form action="#" method="post" name="bar">
				<input type="text" name="in" />
				<input type="submit" value="submit"/>
			</form>
		</body>
	</html>
`

type HelloHandler struct {

}


func (* HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
			io.WriteString(w, form)
	case "POST":
			io.WriteString(w, request.FormValue("in"))
	}
}

func StartHelloWithHttp() {
	//var h HelloHandler
	//http.HandleFunc("/t1", FormServer)
	//http.ListenAndServe(":8000", nil)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        &HelloHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
