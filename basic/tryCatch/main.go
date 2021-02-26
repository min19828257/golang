package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("this is catch")
	}()

	fmt.Println("hello world")

	panic("패닉패닉")

}
