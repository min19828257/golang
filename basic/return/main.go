package main

import "fmt"

func area(w, h int) int {
	return w * h
}

func main() {
	v := area(3, 4)
	fmt.Println(v)
}

// 반환값이 두개 이상일때
// func multiply(w, h int) (int, int) {
// 	return w * 2, h * 2
// }

// func main() {
// 	w, h := multiply(3, 4)
// 	fmt.Println(w, h)
// }
