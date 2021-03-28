package main

import "fmt"

// 포인터(*)를 이용하여 메모리 주소의 값을 변경할 수 있다.
// a의 메모리 주소에 새로운 값을 대입했고, 이는 곧 a에 대입한 것으로 나타난다.
func main() {
	a := 2
	b := &a
	*b = 20
	fmt.Println(a, *b)
	// 결과 20 20
}
