package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var answer []int
	for i, num_1 := range nums {
		for j, num_2 := range nums[i+1:] {
			if num_1+num_2 == target {
				answer = append(answer, i)
				answer = append(answer, i+j+1)
			}
		}
	}
	return answer
}

func main() {
	nums := []int{3, 2, 4}

	fmt.Println(twoSum(nums, 6))
}
