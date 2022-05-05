package main

import "fmt"

func main() {
	var avg float32
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	avg = getAverage(nums, 9)

	fmt.Println(avg)
}

func getAverage(nums []int, size int) float32 {
	var i, sum int
	var average float32

	for i = 0; i < size; i++ {
		sum += nums[i]
	}

	average = float32(sum / size)

	return average

}
