package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		upperA := 'A'
		upperZ := 'Z'
		lowerA := 'a'
		lowerZ := 'z'
		for _, c := range line {
			if upperA <= c && c <= upperZ {
				fmt.Print(strings.ToLower(string(c)))
			} else if lowerA <= c && c <= lowerZ {
				fmt.Print(strings.ToUpper(string(c)))
			} else {
				fmt.Print(string(c))
			}
		}
	}
	fmt.Print("\n")
}
