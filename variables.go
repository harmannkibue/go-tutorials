package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go hello("Armara", &wg)
	wg.Wait()
}

func hello(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello ", name)
}
