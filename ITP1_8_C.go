package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	alphabet := make([]int, 26)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		for _, c := range line {
			index := unicode.ToLower(c) - 'a'
			if index >= 0 {
				alphabet[index]++
			}
		}
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%s : %d\n", string(rune('a'+i)), alphabet[i])
	}
}
