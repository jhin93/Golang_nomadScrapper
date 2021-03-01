package main

import "fmt"

// Low-level Programming
// 메모리 주소를 보는 법, 무엇인가를 메모리 주소에 접근시키는 방법.
// 코드를 작성하다 보면 복사가 되는 것을 원치 않을 수 있다.
// 메모리까지 찾아가고 싶을 수 있음.
// (값 복사의 예시 -> a := 2 // b := a)

func main() {
	a := 2
	b := a
	// 이 시점 이후로 a에 어떤 값이 들어가도 b에 영향이 가지 않는다. 이 시점의 b는 2.
	a = 10
	fmt.Println(a, b)
}
