package main

import "fmt"

func main() {
	var r, c int
	fmt.Scanln(&r, &c)
	A := make([][]int, r+1)
	for i := 0; i < r; i++ {
		A[i] = make([]int, c+1)
		for j := 0; j < c; j++ {
			fmt.Scan(&A[i][j])
		}
	}
	A[r] = make([]int, c+1)
	for i := 0; i < r; i++ {
		sum := 0
		for j := 0; j < c; j++ {
			sum += A[i][j]
		}
		A[i][c] = sum
	}
	for i := 0; i < c+1; i++ {
		sum := 0
		for j := 0; j < r; j++ {
			sum += A[j][i]
		}
		A[r][i] = sum
	}
	for i := 0; i < r+1; i++ {
		for j := 0; j < c+1; j++ {
			fmt.Print(A[i][j])
			if j != c {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
