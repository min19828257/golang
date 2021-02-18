package main

import (
	"fmt"
)

func num(a, b int) <-chan int {
	out := make(chan int)

	go func() {
		//fmt.Println("sub 루틴 시작")
		for i := 0; i < 3; i++ {
			out <- a
			out <- b
			fmt.Println("진행중")
		}
		fmt.Println("진행끝")
		close(out)
	}()

	//time.Sleep(1 * time.Second) // 1초 대기

	return out
}

func sum(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		//fmt.Println("sub2 루틴 시작")
		r := 0
		for i := range c { // wait
			fmt.Println("8-1--", i)
			r = r + i
			fmt.Println("8-2--", i)
		}
		fmt.Println("sub2 진행끝")
		out <- r
	}()

	return out
}

func main() {
	fmt.Println("a")
	c := num(1, 2)
	fmt.Println("b")
	out := sum(c)
	fmt.Println("c")

	for i := 0; i < 5000; i++ {
		fmt.Println("메인이다")
		//time.Sleep(1 * time.Second)
	}

	fmt.Println(<-out)
	fmt.Println("D---")
}
