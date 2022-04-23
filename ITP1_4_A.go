package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scanln(&a, &b)
	fmt.Printf("%d %d %f", a/b, a%b, float64(a)/float64(b))
}
