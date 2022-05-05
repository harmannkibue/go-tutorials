package main

import "fmt"

func main() {
	var a int = 10
	var b *int

	b = &a
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(*b)
}
