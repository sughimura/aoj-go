package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var rooms [4][3][10]int
	for i := 0; i < n; i++ {
		var b, f, r, v int
		fmt.Scanln(&b, &f, &r, &v)
		rooms[b-1][f-1][r-1] += v
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 10; k++ {
				fmt.Printf(" %d", rooms[i][j][k])
				if k == 9 {
					fmt.Print("\n")
				}
			}
			if j == 2 && i != 3 {
				fmt.Print("####################\n")
			}
		}
	}
}
