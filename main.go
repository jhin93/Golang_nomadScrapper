package main

import "fmt"

// Slices.
// 길이를 명시하지 않아도 되는 Array이다.

func main() {
	names := []string{"a", "b", "c"}
	// slice에 item을 추가할 때 사용하는 것이 append()라는 함수이다.
	// append는 2개의 arguments를 사용한다. 첫번째론 해당 slice, 두번째는 넣을 요소
	names = append(names, "d")
	fmt.Println(names)
	// 결과 [a b c d]
}
