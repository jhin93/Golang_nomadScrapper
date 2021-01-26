package main

import (
	"fmt"
)

// fmt는 go에서 표준출력 및 입력을 하기 위한 패키지이다.
// https://brownbears.tistory.com/175

// 반복하는 함수 repeatMe.
//words라는 인자(type은 string)를 받음.
// 원하는 만큼의 arguments를 전달하는 방법은 ...를 찍는 것.
func repeatMe(words ...string) {
	fmt.Println(words)
}

// repeatMe에 여러 개의 arguments를 전달해 호출.
func main() {
	repeatMe("nico", "lynn", "dal", "marl", "flynn")
}
