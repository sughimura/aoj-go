package main

import "fmt"

func main() {
	var n, m, l int
	fmt.Scanln(&n, &m, &l)
	A := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m)
	}
	B := make([][]int, m)
	for i := 0; i < m; i++ {
		B[i] = make([]int, l)
	}
	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, l)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&A[i][j])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < l; j++ {
			fmt.Scan(&B[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			sum := 0
			for k := 0; k < m; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			fmt.Print(C[i][j])
			if j == (l - 1) {
				fmt.Print("\n")
			} else {
				fmt.Print(" ")
			}
		}
	}
}
