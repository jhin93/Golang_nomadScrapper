package main

import "fmt"

// 자바스크립트의 switch와 비슷하다
// if-else가 난무하는 상황을 피할 수 있다.
// switch를 작성하고 나서 if 처럼 variable을 만들 수 있다.

func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(10))
}
