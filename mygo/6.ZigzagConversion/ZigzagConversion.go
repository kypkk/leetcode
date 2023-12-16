package main

import "fmt"

func convert(s string, numRows int) string {

	/* method 1 TLE
	 * if numRows > 1
	 * 將每個 character 依照 zigzag 的順序放入一個 1000 * 1000 的二維 slice
	 * 將 slice 中不為 "" 的 character append to answer (TLE in this step)
	 * if numRows equals to 1, then answer is s
	 */
	// var answer_list [1000][1000]string
	// idx_row := 0
	// idx_col := 0
	// direction := 0
	// answer := ""
	// if numRows > 1 {
	// 	for i := range s {

	// 		if idx_row == numRows {
	// 			direction = 1
	// 			idx_row -= 2
	// 			idx_col += 1
	// 		} else if idx_row == 0 {
	// 			direction = 0
	// 		} else if idx_row < 0 {
	// 			direction = 0
	// 			idx_row = 1
	// 		}
	// 		if direction == 0 {
	// 			answer_list[idx_row][idx_col] = string(s[i])
	// 			idx_row += 1
	// 		} else if direction == 1 {
	// 			answer_list[idx_row][idx_col] = string(s[i])
	// 			idx_row = idx_row - 1
	// 			idx_col = idx_col + 1
	// 		}

	// 	}
	// 	for i := range answer_list {
	// 		for j := range answer_list[i] {
	// 			if answer_list[i][j] != "" {
	// 				answer += string(answer_list[i][j])
	// 			}
	// 		}
	// 	}
	// } else {
	// 	answer = s
	// }

	/* method 2: find the rule for the answer
		* if numRows > 1
		* then loop for numRows times
		* in the ith loop append the ith row of the character to the answer
		* for the first row and the last row of the zigzag, the difference of index between each character is (numrows - 1) * 2 ex: numRows = 3 diff = (3 - 1) * 2 = 4
		* for the row except the first and the last row of the zigzag, the difference of index between each character is ((numrows - 1) * 2 - (the ith row * 2), (the ith row * 2)) ex: the secod row for numRows = 3 diff1 = (3 - 1) * 2 - 1 * 2 = 4 - 2 = 2, diff2 = 2
		* visual diagram for numRows = 3:
		* 0   4
	    * 1 3 5
		* 2   6
	    * visual diagram for numRows = 4:
		* 0     6 diff = 6
		* 1   5 7 diff = (6 - 2, 2)
		* 2 4   8 diff = (6 - 4, 4)
		* 3     9 diff = 6
		* if numRows equals to 1, then answer is s
	*/
	answer := ""
	if numRows > 1 {
		for i := 0; i < numRows; i++ {
			diff_1 := (numRows-1)*2 - i*2
			diff_2 := i * 2
			idx := i
			toggle := 0
			for idx < len(s) {
				answer += string(s[idx])
				if i == 0 {
					idx += diff_1
				} else if i == numRows-1 {
					idx += diff_2
				} else if toggle == 0 {
					idx += diff_1
					toggle = 1
				} else if toggle == 1 {
					idx += diff_2
					toggle = 0
				}
			}
		}
	} else {
		answer = s
	}

	return answer
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 1))
}
