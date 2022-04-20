package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 3)
	for i, _ := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	fmt.Println(a[0], a[1], a[2])
}
