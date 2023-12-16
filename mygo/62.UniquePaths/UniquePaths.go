package main

import (
	"fmt"
	"math"
)

func uniquePaths(m int, n int) int {
	var ret float64
	ret = 1
	x := m + n - 2
	y := int(math.Max(float64(m), float64(n))) - 1
	for i := y; i < x; i++ {
		ret = ret * (float64(i) + 1)
		ret = ret / float64(i+1-y)
	}
	return int(ret)
}

func main() {
	fmt.Println(uniquePaths(16, 16))
}
