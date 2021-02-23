package main

import "fmt"

// if를 사용할 때 ()를 사용하지 않는다.
// if를 쓰는 순간에 variable을 작성할 수 있다. ex) if koreanAge :=

// [1]과 [2]의 작성은 동일하다.

//[1]
//	if koreanAge := age + 2; koreanAge < 18

//[2]
//	koreanAge := age + 2;
//	if koreanAge < 18

// 하지만 [1]의 방식으로 사용하는 것을 추천한다. 그래야 변수가 if-else 구문에서만 사용하려고 만든 것임을 알 수 있다.
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
