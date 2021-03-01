package main

import "fmt"

// Low-Level Programming을 하는 방법 정리.

// 1. 만약 무엇인가의 주소를 보고 싶다면 '&'를 사용하라. &는 주소같은 것.
// ex) a:= 2, fmt.Println(&a)

// 2. 주소를 저장하고 싶다면? 저장할 변수에 &를 이용하여 대입해라.
// ex) b := &a

// 3. 어떤 메모리 주소의 값을 살펴보고 싶다면? *를 이용해라. *는 살펴보는 것.
// ex) b := &a, fmt.Println(*b)

// 4. 어떤 메모리 주소의 값을 변경하고 싶다면? *'변수'에 변경할 값을 대입해라
// ex) *b = 20

// a의 값을 pointer b를 사용해서 변경할 수 있다. *와 &를 사용해서.

func main() {
	a := 2
	b := &a
	*b = 202020
	fmt.Println(a)
}
