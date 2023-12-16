package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	current := 0
	answer := 0
	sub := ""
	for _, c := range s {
		for strings.Contains(sub, string(c)) {
			if current > answer {
				answer = current
			}
			idx := strings.Index(sub, string(c))
			sub = sub[idx+1:]
			current = len(sub)
		}
		sub = sub + string(c)
		current += 1
	}
	if current > answer {
		answer = current
	}
	return answer
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
