package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var num int
	var results int

	fmt.Print("Enter the factorial number: ")
	fmt.Scanln(&input)

	num, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Error occured in string to integer conversion")
	}

	fmt.Println(num)
	results = factorial(num)
	fmt.Printf("The factorial of %s is %d ", input, results)
}

func factorial(input int) int {
	if input <= 1 {
		return 1
	}
	return input * factorial(input-1)
}
