package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// (1)
	var container = make(map[int]int)

	// (2)
	for i, num := range nums {

		// (3)
		complement := target - num

		// (4)
		val, ok := container[complement]
		if ok {
			// (4.b)
			return []int{val, i}
		}

		// (4.a)
		container[num] = i

	}

	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6

	result := twoSum(nums, target)
	fmt.Print(result)
}
