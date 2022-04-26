package main

import "fmt"

func printResult(mark string, data [13]bool) {
	for i, b := range data {
		if b == false {
			fmt.Printf("%s %d\n", mark, i+1)
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	var S, H, C, D [13]bool
	for i := 0; i < n; i++ {
		var mark string
		var number int
		fmt.Scanln(&mark, &number)
		switch mark {
		case "S":
			S[number-1] = true
		case "H":
			H[number-1] = true
		case "C":
			C[number-1] = true
		case "D":
			D[number-1] = true
		}
	}
	printResult("S", S)
	printResult("H", H)
	printResult("C", C)
	printResult("D", D)
}
