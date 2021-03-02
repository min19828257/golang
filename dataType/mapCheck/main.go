package main

import "fmt"

func hasDupeRune(Datamap map[string]string, s string) bool {

	if _, exists := Datamap[s]; exists {
		return true
	}
	return false
}

func main() {

	runeSet := map[string]string{}
	test := "가나다라"
	fmt.Println(hasDupeRune(runeSet, test))
	runeSet["가나다라가"] = "abc"
	test1 := "가나다라가"
	fmt.Println(hasDupeRune(runeSet, test1))
}
