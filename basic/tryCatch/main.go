package main

import "fmt"

func main() {
	defer func() {
		s := recover()
		if s != nil {
			fmt.Println("error :", s) // 에러일때 nil 반환
		} else {
			fmt.Println("success") // 에러가 아닐때 nil 반환
		}
	}()

	var a int = 1
	if a == 1 {
		panic("에러")
	}

	fmt.Println("hello world")
}
