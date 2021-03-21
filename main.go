package main

import "fmt"

func superAdd(fff ...int) int {
	sum := 0
	for a := range fff {
		sum += a
	}
	return sum // 리턴값을 int로 설정했으니 아무 int나 넣어놓은 것.
}

func main() {
	result := superAdd(10, 20, 30)
	fmt.Println(result)
}

// 결과
// 3
// 인덱스를 전부 더한 결과이기에 3이 출력된다. (0+1+2)
// 만약 밸류들을 더하고 싶다면 _, b와 같이 작성해서 인덱스를 무시하고 뒤에 따라오는 밸류만 사용한다.
