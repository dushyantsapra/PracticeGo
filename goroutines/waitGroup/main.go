package main

import (
	"fmt"
	"sync"
	"time"
)

func countDown(n int) {
	for n > 0 {
		fmt.Println("countDown, Value is ", n)
		n--
		time.Sleep(500 * time.Millisecond)
	}
}

func countDownWithWaitGroup(n int, wg *sync.WaitGroup) {
	for n > 0 {
		fmt.Println("countDownWithWaitGroup, Value is ", n)
		n--
		time.Sleep(250 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		countDown(10)
		wg.Done()
	}()

	go countDownWithWaitGroup(5, &wg)
	wg.Wait()
}
