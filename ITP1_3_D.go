package main

import "fmt"

func main() {
	a, b, c := 0, 0, 0
	fmt.Scanln(&a, &b, &c)
	count := 0
	for i := a; i <= b; i++ {
		if c%i == 0 {
			count++
		}
	}
	fmt.Println(count)
}
