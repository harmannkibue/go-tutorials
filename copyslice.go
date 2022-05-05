package main

import "fmt"

func main() {
	var numbers []int
	printSlices(numbers)

	/* append allows nil slice */
	numbers = append(numbers, 0)
	printSlices(numbers)

	/* add one element to slice*/
	numbers = append(numbers, 1)
	printSlices(numbers)

	/* add more than one element at a time*/
	numbers = append(numbers, 2, 3, 4)
	printSlices(numbers)

	/* create a slice numbers1 with double the capacity of earlier slice*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* copy content of numbers to numbers1 */
	copy(numbers1, numbers)
	printSlices(numbers1)
}

func printSlices(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
