package main

import "fmt"

type Number struct {
	A int
	B int
}

func Add(a, b int) int {
	return a + b
}

func main() {
	number := Number{A: 1, B: 2}
	result := Add(number.A, number.B)

	fmt.Println(result)
}
