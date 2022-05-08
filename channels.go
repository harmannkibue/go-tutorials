package main

import "fmt"

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
		for i := 0; i <= 100; i++ {
			data_channel <- i
		}
		close(data_channel)
	}()

	for data := range data_channel {
		fmt.Println(data)
	}

}
