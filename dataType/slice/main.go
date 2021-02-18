package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{3, 4, 5, 6, 7, 8, 32, 4}
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s, "=", s[:3], s[3:5], s[5:])

	// slice 추가
	ns1 := []int{1, 2, 3}
	ns2 := []int{6, 7, 8}
	ns3 := []int{8, 9, 10, 11}

	ns1 = append(ns1, 4, 5, 6)
	fmt.Println(ns1)

	ns2 = append(ns2, ns1...)
	fmt.Println(ns2)

	ns3 = append(ns3, ns2...)
	fmt.Println(ns3)

	//슬라이스 삽입
	s1 := []int{1, 2, 3, 4, 5}

	s1 = insert(s1, []int{-3, -2}, 0) // s: [-3 -2 1 2 3 4 5]
	fmt.Println(s1)

	s1 = insert(s1, []int{0}, 2) // s: [-3 -2 0 1 2 3 4 5]
	fmt.Println(s1)

	s1 = insert(s1, []int{6, 7}, len(s1)) // s: [-3 -2 0 1 2 3 4 5 6 7]
	fmt.Println(s1)

	// 오름차순으로 정렬
	sli := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(sli)
	fmt.Println(sli)
}

func insert(s, new []int, index int) []int {
	return append(s[:index], append(new, s[index:]...)...)
}
