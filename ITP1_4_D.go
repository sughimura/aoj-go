package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	min := a[0]
	max := a[len(a)-1]
	sum := 0
	for _, x := range a {
		sum += x
	}
	fmt.Println(min, max, sum)
}
