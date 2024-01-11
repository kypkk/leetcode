package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	char_map := map[string]string{
		"e": "q", "d": "a", "c": "z", "r": "w", "f": "s", "v": "x", "t": "e", "g": "d", "b": "c", "y": "r", "h": "f", "n": "v", "u": "t", "j": "g", "m": "b", "i": "y", "k": "h", ",": "n", "o": "u",
		"l": "j", ".": "m", "p": "i", ";": "k", "/": ",", "[": "o", "'": "l", "]": "p", "2": "`", "3": "1", "4": "2", "5": "3", "6": "4", "7": "5", "8": "6", "9": "7", "0": "8", "-": "9", "=": "0", "\\": "[", " ": " ",
	}
	ans_str := ""
	for {
		Text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		if Text == "stop\n" {
			break
		}

		for _, char := range Text {
			ans_str += fmt.Sprintf("%s", char_map[string(char)])
		}
		ans_str += fmt.Sprintf("\n")
	}
	fmt.Printf("%s", ans_str)
}
