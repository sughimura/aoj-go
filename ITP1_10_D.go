package main

import (
	"fmt"
	"math"
)

func minkowski(x []float64, y []float64, p float64, n int) (result float64) {
	for i := 0; i < n; i++ {
		result += math.Pow(math.Abs(x[i]-y[i]), p)
	}
	result = math.Pow(result, 1.0/p)
	return
}

func chebyshev(x []float64, y []float64, n int) (result float64) {
	for i := 0; i < n; i++ {
		temp := math.Abs(x[i] - y[i])
		if result < temp {
			result = temp
		}
	}
	return
}

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&y[i])
	}
	fmt.Printf("%f\n", minkowski(x, y, 1.0, n))
	fmt.Printf("%f\n", minkowski(x, y, 2.0, n))
	fmt.Printf("%f\n", minkowski(x, y, 3.0, n))
	fmt.Printf("%f\n", chebyshev(x, y, n))
}
