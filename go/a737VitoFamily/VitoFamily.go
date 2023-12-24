package main

import (
	"fmt"
	"math"
)

func bubbleSort(arr []int) int {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return 0
}

func main() {
	var lines int
	fmt.Scanf("%d", &lines)
	for i := 0; i < lines; i++ {
		var length int
		var arr []int
		answer := 0
		fmt.Scanf("%d", &length)
		for j := 0; j < length; j++ {
			var num int
			fmt.Scanf("%d", &num)
			arr = append(arr, num)
		}
		bubbleSort(arr)
		x := int(len(arr) / 2)
		x = arr[x]
		for j := 0; j < length; j++ {
			answer += int(math.Abs(float64(arr[j] - x)))
		}
		fmt.Println(answer)
	}
}
