package main

import "fmt"

// 2020.3.13 지금까지 한것에 대한 복습.

// 함수에서 리턴을 여러개 해보기. 타입이 서로 다른 리턴값을 여러개 만들 순 없을까?
// https://fenderist.tistory.com/187
func multiply() (int, string) {

	return 1, "stir"
}

func main() {
	a, b := multiply()
	fmt.Println(a, b)
}
