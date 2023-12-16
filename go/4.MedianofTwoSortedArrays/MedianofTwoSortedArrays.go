package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var answer_arr []int
	answer := 0.00000
	total_length := len(nums1) + len(nums2)
	idx_1 := 0
	idx_2 := 0
	for i := 0; i < total_length; i++ {
		if idx_1 != len(nums1) && idx_2 != len(nums2) {
			if nums1[idx_1] <= nums2[idx_2] {
				answer_arr = append(answer_arr, nums1[idx_1])
				idx_1 += 1
			} else if nums2[idx_2] < nums1[idx_1] {
				answer_arr = append(answer_arr, nums2[idx_2])
				idx_2 += 1
			}
		} else if idx_1 == len(nums1) {
			answer_arr = append(answer_arr, nums2[idx_2])
			idx_2 += 1
		} else {
			answer_arr = append(answer_arr, nums1[idx_1])
			idx_1 += 1
		}
	}
	if total_length%2 == 0 {
		answer = float64(answer_arr[total_length/2-1]+answer_arr[total_length/2]) / 2
	} else {
		answer = float64(answer_arr[int(total_length/2)])
	}
	return answer

}

func main() {
	list_1 := []int{1, 3}
	list_2 := []int{2, 4}
	fmt.Println(findMedianSortedArrays(list_1, list_2))
}
