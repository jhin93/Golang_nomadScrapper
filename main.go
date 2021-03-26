package main

import "fmt"

// 메모리.

func main() {
	a := 2
	b := a
	a = 10
	// 이 시점 이후로 a에 어떤 값이 들어가도 b에 영향이 가지 않는다. 이 시점의 b는 2.
	fmt.Println(a, b)

}
