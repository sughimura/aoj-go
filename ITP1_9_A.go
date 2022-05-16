package main

import (
	"fmt"
	"strings"
)

func main() {
	var w string
	var count int
	fmt.Scan(&w)
	for {
		var t string
		fmt.Scan(&t)
		if t == "END_OF_TEXT" {
			break
		}
		if strings.ToLower(t) == w {
			count++
		}
	}
	fmt.Println(count)
}
