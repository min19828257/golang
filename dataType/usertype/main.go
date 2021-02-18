package main

import "fmt"

type quantity int
type dozen []quantity

func main() {

	var d dozen
	for i := quantity(1); i <= 12; i++ {
		d = append(d, i)
	}
	fmt.Println(d)

}
