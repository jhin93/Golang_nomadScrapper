package main

import (
	"fmt"
)

func superAdd(fff ...int) int {
	for a, b := range fff {
		fmt.Println(a, b)
	}
	return 2 // 리턴값을 int로 설정했으니 아무 int나 넣어놓은 것.
}

func main() {
	superAdd(10, 20, 30)
}

// 결과
// 0 10
// 1 20
// 2 30
