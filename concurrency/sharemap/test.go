package main

import "fmt"

// 범용적인 맵
func main() {
	abc := make(map[string]interface{})
	abc["123"] = 1
	abc["1243"] = '1'
	abc["1243"] = "abcd"
	fmt.Println(abc)

}
