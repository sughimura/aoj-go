package main

import "fmt"

func main() {
	var s, p string
	fmt.Scanln(&s)
	fmt.Scanln(&p)
	var ss = s + s
	for i := 0; i < len(s); i++ {
		if p == ss[i:i+len(p)] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
