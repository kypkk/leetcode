package main

import (
	"fmt"
)

func main() {
	for {
		var num1, num2 int
		fmt.Scanf("%d%d", &num1, &num2)
		if num1 == 0 && num2 == 0 {
			break
		}
		carries := 0
		carry := 0
		for num1 != 0 && num2 != 0 {
			a := num1 % 10
			b := num2 % 10
			if a+b+carry >= 10 {
				carries += 1
				carry = 1
			} else {
				carry = 0
			}
			num1 = int((num1 - a) / 10)
			num2 = int((num2 - b) / 10)
		}
		if carries == 0 {
			fmt.Printf("No carry operation.\n")
		} else if carries == 1 {
			fmt.Printf("1 carry operation.\n")
		} else {
			fmt.Printf("%d carry operations.\n", carries)
		}
	}
}
