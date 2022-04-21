package main

import (
	"fmt"
	"sort"
)

func main() {
	for {
		x, y := 0, 0
		fmt.Scanln(&x, &y)
		if x == 0 && y == 0 {
			break
		} else {
			a := []int{x, y}
			sort.Ints(a)
			fmt.Println(a[0], a[1])
		}
	}
}
