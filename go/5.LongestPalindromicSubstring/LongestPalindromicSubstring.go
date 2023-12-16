package main

import "fmt"

func longestPalindrome(s string) string {
	// first step: initialize the dp table with len(s) * len(s) all false
	length := len(s)
	answer := ""
	max := 0
	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, length)
	}
	/* second step:
	   a. Set all dp [i][i] to true, since all the single charater are palindrome
	*/
	for i := range dp {
		dp[i][i] = true
		if i < length-1 {
			if s[i] == s[i+1] {
				dp[i+1][i] = true
				dp[i][i+1] = true
			}
		}
	}

	/* third step:
	   a. For all dp[i][j] check if the inner substring is already a palindrome
	   b. If the inner substring(inner substring in the dp table: dp[i+1][j-1]) is a palindrome and the i's and j's character are the same then s[i:j+1] is palindrome
	   c. compare the length with the max length palindrome
	*/
	for i := range dp {
		for j := range dp[i] {
			if i+1 < j {
				dp[i][j] = dp[i+1][j-1] && (s[i] == s[j])
				dp[j][i] = dp[i+1][j-1] && (s[i] == s[j])
			}
			if j+1 < i {
				dp[j][i] = dp[i-1][j+1] && (s[j] == s[i])
				dp[i][j] = dp[i-1][j+1] && (s[i] == s[j])
			}

			if dp[i][j] || dp[j][i] {
				if max < j+1-i {
					max = j + 1 - i
					answer = s[i : j+1]
				} else if max < i+1-j {
					max = i + 1 - j
					answer = s[j : i+1]
				}
			}

		}
	}
	return answer
}

func main() {
	fmt.Println(longestPalindrome("absba"))
}

//   a b c b a
// a t f f f f
// b f t f t f
// c f f t f f
// b f t f t f
// a f f f f t
