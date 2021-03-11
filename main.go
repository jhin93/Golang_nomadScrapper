package main

import "fmt"

type test struct {
	aa int
	bb int
}

func main() {
	a := test{aa: 2, bb: 4}
	b := &a
	fmt.Println(a, b)

	// 결과 : {2 4} &{2 4}
}
