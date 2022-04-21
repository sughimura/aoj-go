package main

import "fmt"

func main() {
	i := 1
	for {
		x := 0
		fmt.Scanln(&x)
		if x == 0 {
			break
		} else {
			fmt.Printf("Case %d: %d\n", i, x)
		}
		i++
	}
}
