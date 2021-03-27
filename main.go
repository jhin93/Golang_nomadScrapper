package main

import "fmt"

// 만약 다음과 같이 b에 a의 값을 담으면 이후에 a의 값이 변경되어도 b의 값은 변경되지 않는다.
func main() {
	a := 2
	b := a
	a = 5
	fmt.Println(b)
	// 결과 2
}
