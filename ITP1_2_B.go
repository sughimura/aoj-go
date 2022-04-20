package main

import "fmt"

func main() {
	a, b, c := 0, 0, 0
	fmt.Scanln(&a, &b, &c)
	if a < b && b < c {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
