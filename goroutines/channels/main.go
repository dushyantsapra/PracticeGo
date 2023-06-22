package main

import "fmt"

func fibonacci(n int, ch chan int) {
	fib := make([]int, n)
	fib[0] = 0
	ch <- fib[0]
	fib[1] = 1
	ch <- fib[1]
	for iLoop := 2; iLoop < n; iLoop++ {
		fib[iLoop] = fib[iLoop-1] + fib[iLoop-2]
		ch <- fib[iLoop]
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go fibonacci(10, ch)

	for value := range ch {
		fmt.Println("Value is ", value)
	}
}
