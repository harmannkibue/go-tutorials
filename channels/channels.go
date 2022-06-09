package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork(i int) int {
	time.Sleep(time.Second)
	return rand.Intn(1000)
}

func main() {
	data_channel := make(chan int)

	////use go routines
	//go func() {
	//	data_channel <- 789
	//}()
	//
	//data_received := <-data_channel
	//
	//fmt.Println(data_received)

	//Data channel should be able to read and write concurrently
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i <= 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := doWork(i)
				data_channel <- result
			}()
		}
		wg.Wait()
		defer close(data_channel)
	}()

	for data := range data_channel {
		fmt.Println(data)
	}
}
