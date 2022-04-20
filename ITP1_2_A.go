package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scanln(&a, &b)
	if a < b {
		fmt.Println("a < b")
	} else if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a == b")
	}
}
