package main

import "fmt"

func main() {
	numberMap := map[string]int{}
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3
	fmt.Println(numberMap)
	//map[one:1, two:2, three:3]
}
