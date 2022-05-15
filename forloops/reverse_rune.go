package main

import "fmt"

func main() {
	run := "123456789"

	for _, ch := range run {
		defer fmt.Printf("%c", ch)
	}
}
