package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var inputNum string
		fmt.Scanf("%s", &inputNum)
		if inputNum == "0" {
			break
		}
		sum := 0
		op := 0
		for _, ch := range inputNum {
			if op == 0 {
				value := ch - '0'
				sum += int(value)
				op = 1
			} else {
				value := ch - '0'
				sum -= int(value)
				op = 0
			}
		}
		sum = int(math.Abs(float64(sum)))
		if sum%11 == 0 {
			fmt.Printf("%s is a multiple of 11\n", inputNum)
		} else {
			fmt.Printf("%s is not a multiple of 11\n", inputNum)
		}
	}
}
