package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			return
		}
		s := make([]float64, n)
		for i, _ := range s {
			fmt.Scan(&s[i])
		}
		var sum float64
		for _, value := range s {
			sum += value
		}
		m := (sum * 1.0) / (float64(n) * 1.0)
		var smSum float64
		for _, value := range s {
			smSum += math.Pow(value-m, 2.0)
		}
		a := math.Pow(smSum/float64(n*1.0), 0.5)
		fmt.Printf("%.10f\n", a)
	}
}
