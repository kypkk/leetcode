package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var lines int // lines of input
	fmt.Scanf("%d", &lines)
	ans := make(map[string]int) // the answer map to store all the character and their amounts

	// collecting all the inputs by line and store them into ans map
	for i := 0; i < lines; i++ {
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		for _, char := range line {
			if int(char) >= int('a') && int(char) <= int('z') {
				char = char - 'a' + 'A'
			}
			if int(char) >= int('A') && int(char) <= int('Z') {
				ans[string(char)] += 1
			}
		}
	}

	// The following are the map sorting part
	// Firstly, create a two dimension slice type interface because we want to store the character and the amount together
	pair := make([][2]interface{}, 0, len(ans))
	for key, val := range ans {
		pair = append(pair, [2]interface{}{key, val})
	}

	// Secondly, sort the slice by the index1 which is the amount of the character
	sort.Slice(pair, func(i, j int) bool {
		return pair[i][1].(int) >= pair[j][1].(int)
	})

	// Thirdly, extract the keys of the sorted slice
	keys := make([]string, 0, len(pair))
	for _, p := range pair {
		keys = append(keys, p[0].(string))
	}

	// Lastly, print out the map by the sorted keys
	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, ans[key])
	}
}
