package main

import (
	"fmt"
)

func main() {
	for {
		var num1, num2 int
		fmt.Scanf("%d", &num1)
		fmt.Scanf("%d", &num2)
		fmt.Printf("%d ", num1)
		fmt.Printf("%d ", num2)
		idx := 0
		bound := 0
		max := 0

		if num1 > num2 {
			idx = num2
			bound = num1
		} else {
			idx = num1
			bound = num2
		}

		for i := idx; i < bound; i++ {
			n := i
			sum := 1
			for n != 1 {
				if n%2 == 1 {
					n = 3*n + 1
				} else {
					n = int(n / 2)
				}
				sum += 1
			}
			if sum > max {
				max = sum
			}
		}
		fmt.Printf("%d\n", max)

	}
}
