package main

import "fmt"

func main() {
	a := 2
	b := &a
	fmt.Println(a, b)
	// 2 0xc00010e008
}
