package main

import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // 이런식으로 조건문 안에 변수를 만드는 게 오직 조건문에서만 사용되는 변수라는 걸 알려준다.
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
