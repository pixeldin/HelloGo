package main

import (
	"HelloGo/basic/http/utils"
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	//http.Handle("/foo", fooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		//time.Sleep(100 * time.Millisecond)
		fmt.Printf("%s Request enter...\n", utils.GetPresentFormat())
		fmt.Fprintf(w, "Hello there, from %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
