package main

import "fmt"

type Item struct {
	name     string
	price    float64
	quantity int
}

// Item 타입에 Cost() 메서드 정의
func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func main() {
	// 아이템 값 생성
	shirt := Item{name: "men's Slim-Fit Shirt", price: 25000, quantity: 3}

	//Cost() 메서드로 가격 출력
	fmt.Println(shirt.Cost()) // 75000
}
