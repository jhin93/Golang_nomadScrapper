package main

import "fmt"

// name := "nico" // 함수 밖에서는 사용할 수 없다.
func main() {
	// 축약형은 오직 func 안에서만, 그리고 변수(var)에만 적용 가능하다. 축약형이 존재하면 Go가 첫번째 값을 기준으로 변수의 type을 찾아 정해준다.
	name := "nico" // := 는 type입력을 생략하게 해준다. 정해진 type은 임의로 변경할 수 없다.
	name = "hjie"
	fmt.Println(name)
}
