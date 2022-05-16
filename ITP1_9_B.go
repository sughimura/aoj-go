package main

import "fmt"

func main() {
	for {
		var str string
		fmt.Scan(&str)
		if str == "-" {
			return
		}
		var m int
		fmt.Scan(&m)
		for i := 0; i < m; i++ {
			var h int
			fmt.Scan(&h)
			str = str[h:] + str[:h]
		}
		fmt.Println(str)
	}
}
