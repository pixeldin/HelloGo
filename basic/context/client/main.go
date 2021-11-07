package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()

	// 控制请求超时
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}
}
