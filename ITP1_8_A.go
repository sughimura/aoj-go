package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			if unicode.IsUpper(c) {
				fmt.Print(strings.ToLower(string(c)))
			} else if unicode.IsLower(c) {
				fmt.Print(strings.ToUpper(string(c)))
			} else {
				fmt.Print(string(c))
			}
		}
	}
	fmt.Print("\n")
}
