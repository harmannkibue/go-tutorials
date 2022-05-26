package main

import (
	"fmt"
	"time"
)

// This is a function that returns a channel

func main() {
	c := gen("MESSAGE")
	for i := 0; i < 5; i++ {
		fmt.Println("You are saying ", <-c)
	}
	fmt.Println("You are boring and I am leaving!")
}

func gen(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%d th %s\n", i, msg)
			time.Sleep(time.Second)
		}
	}()
	return c
}
