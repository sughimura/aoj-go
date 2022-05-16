package main

import "fmt"

func main() {
	var n int
	var taroP, hanaP int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var taro, hana string
		fmt.Scanln(&taro, &hana)
		if taro > hana {
			taroP += 3
		} else if taro < hana {
			hanaP += 3
		} else {
			taroP++
			hanaP++
		}
	}
	fmt.Printf("%d %d\n", taroP, hanaP)
}
