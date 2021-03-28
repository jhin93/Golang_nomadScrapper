package main

import "fmt"

// array처럼 map에도 range를 적용할 수 있다.

func main() {
	nico := map[int]string{1: "a", 2: "b", 3: "c"}
	for a, b := range nico {
		fmt.Println(a, b)
	}
	// 결과
	// 1 a
	// 2 b
	// 3 c
	for a := range nico {
		fmt.Println(a)
	}
	// 결과
	// 1
	// 2
	// 3
	for _, b := range nico {
		fmt.Println(b)
	}
	// 결과
	// a
	// b
	// c
}
