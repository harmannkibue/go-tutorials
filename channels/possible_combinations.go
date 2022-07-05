package main

import (
	"fmt"
	"strconv"
)

//func posibleCombinations() []string {
//	// This is code to generate possible codes between 000 - 999
//	var codesList []string
//	for i := 0; i < 10; i++ {
//		for j := 0; j < 10; j++ {
//			for k := 0; k < 10; k++ {
//				code := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k)
//				codesList = append(codesList, code)
//			}
//		}
//	}
//	fmt.Println("THe length of the list ", len(codesList))
//	return codesList
//}

// Returns a list of string codes that fits the given number of digits.example given 3 it will generate from 000 to 999
func rCombinations(p int, n []int, c []int, cc [][]int) [][]int {
	if len(n) == 0 || p <= 0 {
		return cc
	}
	p--
	for i := range n {
		r := make([]int, len(c)+1)
		copy(r, c)
		r[len(r)-1] = n[i]
		if p == 0 {
			cc = append(cc, r)
		}
		cc = rCombinations(p, n, r, cc)
	}
	return cc
}

func Combinations(p int) []string {
	var codesList []string
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := rCombinations(p, numbers, nil, nil)
	for i := range res {
		code := strconv.Itoa(res[i][0]) + strconv.Itoa(res[i][1]) + strconv.Itoa(res[i][2])
		codesList = append(codesList, code)
	}
	return codesList
}

func main() {
	pools := 3
	c := Combinations(pools)

	fmt.Println("pools generated iss ", c)
	//for i := range c {
	//	//fmt.Println("Starting the range ", i)
	//	fmt.Println("The type of gen iss \n", c[i][0])
	//}
}
