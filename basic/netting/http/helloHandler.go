package http

import (
	"fmt"
	"io"
	"net/http"
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

type HelloHandler struct{}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/t1", FormServer)
	http.HandleFunc("/t2", ServeHTTP)
	http.ListenAndServe(":8000", nil)
}
