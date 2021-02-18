package main

import "fmt"

type Option struct {
	name  string
	value string
}

type Item struct {
	name     string
	price    float64
	quantity int
	Option   // 임베디드 필드
}

func main() {
	shoes := Item{"Sports Shoes", 30000, 2, Option{"color", "red"}}

	fmt.Println(shoes)

	// name  필드에 접근
	fmt.Println(shoes.name)

	fmt.Println(shoes.Option.name)
}
