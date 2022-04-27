package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	A := make([][]int, n)
	b := make([]int, m)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&A[i][j])
		}
	}
	for i := 0; i < m; i++ {
		fmt.Scanln(&b[i])
	}
	for i := 0; i < n; i++ {
		result := 0
		for j := 0; j < m; j++ {
			result += A[i][j] * b[j]
		}
		c[i] = result
	}
	for i := 0; i < n; i++ {
		fmt.Println(c[i])
	}
}
