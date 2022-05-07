package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x string
	for {
		fmt.Scanln(&x)
		if x == "0" {
			break
		}
		sum := 0
		for _, c := range x {
			si, _ := strconv.Atoi(string(c))
			sum += si
		}
		fmt.Println(sum)
	}
}
