package main

import (
	"HelloGo/basic/rscCtl/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines = 25
	poolMaxSize   = 10
)

type dbConnection struct {
	ID int32
}

// 实现了 io.Closer 接口，以便 dbConnection
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

var idCounter int32

func CreateConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create new Connection", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	pl, err := pool.New(CreateConnection, poolMaxSize)
	if err != nil {
		log.Println("New pool err: ", err)
	}

	for q := 0; q < maxGoroutines; q++ {
		go func(i int) {
			//连接消耗
			performQueries(i, pl)
			wg.Done()
		}(q)
	}

	wg.Wait()

	log.Println("Shutdown Program.")
	pl.Close()
}

func performQueries(query int, p *pool.Pool) {

	for {

		conn, err := p.Acquire()
		if err != nil {
			log.Println(err)
			break
		}
		if conn != nil {
			//执行完成, 释放连接
			defer func(c io.Closer) {
				p.Release(c)
				log.Printf("Release conn, QID[%d], CID[%d]\n", query, conn.(*dbConnection).ID)
			}(conn)

			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			log.Printf("QID[%d] executed, CID[%d]\n", query, conn.(*dbConnection).ID)
			break
		} else {
			//等待其他占用者释放
			log.Printf("QID[%d], pool dry,"+
				" please wait other's releasing...\n", query)
		}

	}

	return
}
