package main

import "fmt"

func make2dSlice(a int, b int) [][]int {
	matrix := make([][]int, a)
	for i := 0; i < a; i++ {
		matrix[i] = make([]int, b)
	}
	return matrix
}

func main() {
	var n, m, l int
	fmt.Scanln(&n, &m, &l)
	A := make2dSlice(n, m)
	B := make2dSlice(m, l)
	C := make2dSlice(n, l)

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
