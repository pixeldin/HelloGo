package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main()  {
	//resp, err := http.Get(os.Args[1])
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatalln(err)
	}

	//file, err := os.Create(os.Args[2])
	file, err := os.Create("resp.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	//输出源
	dest := io.MultiWriter(os.Stdout, file)

	//Out put response body into
	io.Copy(dest, resp.Body)
	if err := resp.Body.Close(); err != nil {
		log.Println(err)
	}

}
