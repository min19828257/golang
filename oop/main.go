package main

import (
	"fmt"

	polymorphism "github.com/min19828257/oop/polymorph"
)

type Coster interface {
	Cost() float64
}

func displayCost(c Coster) {
	fmt.Println("cost: ", c.Cost())
}

type DiscountItem struct {
	polymorphism.Item
	discountRate float64
}

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

func main() {

	eventShoes := DiscountItem{
		polymorphism.Item{"Women's Walking Shoes", 50000, 3},
		10.00,
	}

	fmt.Println(eventShoes)
	displayCost(eventShoes)
}
