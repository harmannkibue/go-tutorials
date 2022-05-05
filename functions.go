package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var ret int

	ret = maxNumber(a, b)

	fmt.Println(ret)

	var first string = "Harman"
	var second string = "Kibue"

	var j, k string = swapNames(first, second)
	fmt.Println(j, k)

}

func maxNumber(a, b int) int {

	if a > b {
		return a
	} else {
		return b
	}
}

func swapNames(first, second string) (string, string) {
	return second, first
}
