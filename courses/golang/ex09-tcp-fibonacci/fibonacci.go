package main

import "math/big"

type FibonacciRequest struct {
	Number int `json:"number"`
}

type FibonacciResponse struct {
	Result    string `json:"result"`
	Timedelta int64  `json:"timedelta"`
}

func fibonacci(number int) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 0; i < number; i++ {
		a.Add(a, b)
		a, b = b, a
	}

	return a
}
