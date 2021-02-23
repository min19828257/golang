package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt") // hello.txt 파일 생성
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() //main 함수가 끝나기 직전에 파일을 닫음

	s := "hello, world"
	n, err := file.Write([]byte(s))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")
}
