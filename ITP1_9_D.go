package main

import (
	"fmt"
)

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	var str string
	fmt.Scanln(&str)
	var q int
	fmt.Scanln(&q)
	for i := 0; i < q; i++ {
		var cmd, p string
		var a, b int
		fmt.Scan(&cmd, &a, &b)
		if cmd == "replace" {
			fmt.Scanln(&p)
		}
		if cmd == "replace" {
			str = str[:a] + p + str[b+1:]
		} else if cmd == "reverse" {
			str = str[:a] + reverse(str[a:b+1]) + str[b+1:]
		} else {
			fmt.Println(str[a : b+1])
		}
	}
}
