package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var x string
	for {
		fmt.Scanln(&x)
		if x == "0" {
			break
		}
		xList := strings.Split(x, "")
		sum := 0
		for _, s := range xList {
			si, _ := strconv.Atoi(s)
			sum += si
		}
		fmt.Println(sum)
	}
}
