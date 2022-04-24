package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Printf(" %d", i)
			continue
		}
		x := i
		for {
			if x%10 == 3 {
				fmt.Printf(" %d", i)
				break
			}
			x /= 10
			if x == 0 {
				break
			}
		}
	}
}
