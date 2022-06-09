package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	x += 1
	ch <- x
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 1)
	for i := 0; i < 101; i++ {
		wg.Add(1)
		go increment(&wg, ch)
	}
	wg.Wait()
	fmt.Println("THe value of x iss ", x)
}
