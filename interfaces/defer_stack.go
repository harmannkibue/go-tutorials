package main

import "fmt"

func main() {
	name := "Harman"
	fmt.Println("The name before the defer iss ", name)

	//we print the name in reverse order using the defer call
	for _, char := range name {
		defer fmt.Printf("%c", char)
	}
}
