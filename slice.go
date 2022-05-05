package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	numbers[0] = 20
	printSlic(numbers)
}

func printSlic(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
