package main

import "fmt"

// b 또한 주소이기 때문에, b를 이용하여 a의 값을 변경할 수도 있다.
// b는 a를 살펴보는 pointer이다. b := &a
// *b = '새로 대입할 값'.
// 위와 같이 대입하면 a의 메모리 주소와 연결되어있기 때문에 a 메모리 주소의 값이 변경된다.
func main() {
	a := 2
	b := &a
	*b = 20
	fmt.Println(a)
}
