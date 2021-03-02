package main

import "fmt"

func main() {
	sum := 0
	//for 문에 초기화 구문, 조건식, 후속작업 정의
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	names := []string{"홍길동", "이순신", "강감찬"}

	for index, name := range names {
		println(index, name)
	}
}
