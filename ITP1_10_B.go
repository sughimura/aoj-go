package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, C float64
	fmt.Scan(&a, &b, &C)
	S := (1.0 / 2.0) * a * b * math.Sin(C*math.Pi/180.0)
	c := math.Pow(a*a+b*b-2*a*b*math.Cos(C*math.Pi/180.0), 1.0/2.0)
	L := a + b + c
	h := 2 * S / a
	fmt.Println(S)
	fmt.Println(L)
	fmt.Println(h)
}
