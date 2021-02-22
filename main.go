package main

import "fmt"

// Go에서 loop는 오로지 for를 사용하는 것으로만 가능하다.
// range는 array에 loop를 적용할 수 있도록 해준다.
func superAdd(numbers ...int) int {
	for a, b := range numbers {
		fmt.Println(a, b)
	} // a 에는 index가, b에는 값이 출력된다.
	return 2 // return type이 int이라서 return 1 을 하는 것.
}

func main() {
	superAdd(1, 2, 3, 4, 5, 6)
}

// 결과
// 0 1
// 1 2
// 2 3
// 3 4
// 4 5
// 5 6
