package main

import "fmt"

type Duck interface {
	Quack() string // 쾍쾍 거리는 동작
	Walk() string  // 걷는 동작
}

type Swan struct {
}

func (s Swan) Quack() string {
	return "고니 울름소리"
}
func (s Swan) Walk() string {
	return "고니 걷는 모습"
}

type Chicken struct {
}

func (c Chicken) Quack() string {
	return "닭 울름소리"
}
func (c Chicken) Walk() string {
	return "닭 걷는 모습"
}

//깜짝 놀랬을 때 울고 난 후 걷는다(도망) 간다
func Scare(duck Duck) {
	// Duck의 동작을 구현한 어떤 객체라도 argument로 전달 될 수 있다.
	fmt.Println(duck.Quack())
	fmt.Println(duck.Walk())
}

func main() {
	var swan Duck = Swan{}
	Scare(swan)
	var chicken Duck = Chicken{}
	Scare(chicken)
}
