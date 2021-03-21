package main

import "fmt"

func superAdd(fff ...int) int {
	sum := 0
	for _, b := range fff {
		sum += b
	}
	return sum // 리턴값을 int로 설정했으니 아무 int나 넣어놓은 것.
}

func main() {
	result := superAdd(10, 20, 30)
	fmt.Println(result)
}

// 결과
// 60
