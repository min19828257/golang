package main

import "fmt"

type Item struct {
	name     string
	price    int64
	quantity int64
}

func main() {
	shirt := Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	fmt.Println(shirt)

	p := &Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	fmt.Println(p)

	item := new(Item)
	item.name = "Men's Slim-Fit Shirt"
	item.price = 25000
	item.quantity = 3

	fmt.Println(item)
	fmt.Println(item.price)
}
