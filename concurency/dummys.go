package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go boring("Borring", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring I am leaving!")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s at %d", msg, i)
		time.Sleep(time.Second)
	}
}
