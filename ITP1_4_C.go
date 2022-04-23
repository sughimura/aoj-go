package main

import "fmt"

func main() {
L:
	for {
		var a, b int
		var op string
		fmt.Scan(&a, &op, &b)
		result := 0
		switch op {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			result = a / b
		default:
			break L
		}
		fmt.Println(result)
	}
}
