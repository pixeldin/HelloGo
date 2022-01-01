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

	hf := func(w http.ResponseWriter, r *http.Request) {
		//time.Sleep(100 * time.Millisecond)
		fmt.Printf("%s Request enter...\n", utils.GetPresentFormat())
		fmt.Fprintf(w, "Hello there, from %q", html.EscapeString(r.URL.Path))
	}

	http.HandleFunc("/bar", hf)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
