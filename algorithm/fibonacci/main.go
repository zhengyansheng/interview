package main

import "fmt"

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func main() {
	// fib(0) = 0; fib(1) = 1
	// 0, 1, 1, 2, 3, 5, 8, 13, 21
	s := fib(8)
	fmt.Println(s)
}
