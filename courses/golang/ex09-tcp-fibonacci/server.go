package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
)

type Cache struct {
	lock    sync.RWMutex
	storage map[int]string
}

func main() {
	port := "8081"
	fmt.Println("Server is listening on port " + port)

	ln, _ := net.Listen("tcp", ":"+port)
	defer ln.Close()

	cache := &Cache{
		lock:    sync.RWMutex{},
		storage: make(map[int]string),
	}

	for {
		if conn, err := ln.Accept(); err != nil {
			fmt.Println(err)
			return
		} else {
			go handleConnection(conn, cache)
		}
	}
}

func handleConnection(conn net.Conn, cache *Cache) {
	for {
		decoder := json.NewDecoder(conn)
		var req FibonacciRequest
		err := decoder.Decode(&req)
		if err != nil {
			continue
		}
		fmt.Printf("Number Received: %d\n", req.Number)

		cache.lock.Lock()
		start := time.Now()
		var result string

		if s, ok := cache.storage[req.Number]; ok {
			result = s
		} else {
			fibNumber := fibonacci(req.Number)
			result = fibNumber.String()
			cache.storage[req.Number] = result
		}

		end := time.Now()
		cache.lock.Unlock()

		res, _ := json.Marshal(&FibonacciResponse{
			Result:    result,
			Timedelta: int64(end.Sub(start)),
		})
		conn.Write(res)
	}
	conn.Close()
}
