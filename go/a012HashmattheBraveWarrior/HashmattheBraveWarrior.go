package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var num1, num2 int64
		fmt.Scanf("%d%d", &num1, &num2)
		fmt.Println(int(math.Abs(float64(num2 - num1))))
	}
}
