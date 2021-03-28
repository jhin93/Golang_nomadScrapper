package main

import "fmt"

// Map.
// Object와 비슷한 구조의 데이터 형식이다.

func main() {
	//작성하는 법 - map[key타입]value타입{key:value, key:value, ...}
	nico := map[int]string{1: "a", 2: "b", 3: "c"}
	// 밸류를 작성할 때 작은 따옴표('')가 아닌 큰 따옴표("")로 작성한다. 그렇지 않으면 에러 발생.
	fmt.Println(nico)
}
