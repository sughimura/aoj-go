package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for i := n - 1; i >= 0; i-- {
		fmt.Print(a[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
