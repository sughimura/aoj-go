package main

import "fmt"

func main() {
	for {
		var h, w int
		fmt.Scanln(&h, &w)
		if h == 0 && w == 0 {
			break
		}
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if (i%2 == 0 && j%2 == 0) || (i%2 == 1 && j%2 == 1) {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}
