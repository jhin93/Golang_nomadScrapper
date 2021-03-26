package main

import "fmt"

// 메모리.

func main() {
	a := 2
	fmt.Println(&a, a)
	//결과 : 0xc0000b2008 2
}
