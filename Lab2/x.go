package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	start := time.Now()

	fib31 := fib(31)

	elapsed := time.Since(start)
	fmt.Printf("Fibonacci(30) = %d\n", fib31)
	fmt.Printf("Czas wykonania: %s\n", elapsed)
}