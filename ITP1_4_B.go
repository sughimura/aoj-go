package main

import (
	"fmt"
	"math"
)

func main() {
	r := 0.0
	fmt.Scanln(&r)
	fmt.Printf("%f %f", r*r*math.Pi, 2*r*math.Pi)
}
