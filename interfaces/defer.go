package main

import "fmt"

func finished() {
	fmt.Println("Finished finding largest")
}

func findMaxNumn(nums []int) {
	defer finished()
	fmt.Println("Started finding the largest")
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	fmt.Printf("The largest number in %d is %d \n", nums, max)
}

func main() {
	nums := []int{
		1, 4, 5, 9, 4, 6, 0, 2,
	}

	findMaxNumn(nums)

}
