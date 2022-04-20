package main

import "fmt"

func main() {
	W, H, x, y, r := 0, 0, 0, 0, 0
	fmt.Scanln(&W, &H, &x, &y, &r)
	if 0 <= (x-r) && (x+r) <= W && 0 <= (y-r) && (y+r) <= H {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
