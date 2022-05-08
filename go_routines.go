package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		countThings("Plants")
		wg.Done()
	}()

	wg.Wait()
}

func countThings(thing string) {
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d thing(s)\n", i)
		time.Sleep(time.Millisecond * 500)
	}
}
