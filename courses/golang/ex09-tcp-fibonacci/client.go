package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	port := "8081"
	conn, err := net.Dial("tcp", "127.0.0.1:"+port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		//fmt.Print("Number: ")
		s, _ := reader.ReadString('\n')
		if s == "" {
			return
		}

		number, err := strconv.Atoi(strings.Trim(s, "\n"))
		if err != nil {
			fmt.Println(err)
			continue
		}

		req, err := json.Marshal(&FibonacciRequest{
			Number: number,
		})
		if err != nil {
			fmt.Println(err)
			continue
		}
		conn.Write(req)

		decoder := json.NewDecoder(conn)
		var res FibonacciResponse
		_ = decoder.Decode(&res)

		fmt.Println(res.Timedeltai+"µs", res.Result)
		// fmt.Printf("Fibonacci number: %s\n", res.Result)
		// fmt.Printf("Spent time: %dns\n", res.Timedelta)
	}
}
