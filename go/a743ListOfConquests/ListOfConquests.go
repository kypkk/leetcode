package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var lines int
	fmt.Scanf("%d", &lines)
	ans := make(map[string]int)

	// getting input of each lines
	for i := 0; i < lines; i++ {
		var country string
		fmt.Scanln(&country)
		LoverName, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			_ = fmt.Errorf(" %v error", err)
			return
		}
		LoverName = strings.TrimSuffix(LoverName, "\n")
		ans[country] += 1
	}

	// sort the keys
	var keys []string
	for k := range ans {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// print out the answer by sorted keys
	for _, k := range keys {
		fmt.Printf("%s %d\n", k, ans[k])
	}
}
