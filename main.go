package main

import "fmt"

// Go에서 loop는 오로지 for를 사용하는 것으로만 가능하다.
// range는 array에 loop를 적용할 수 있도록 해준다.
// 또한 range는 오로지 for 안에서만 적용할 수 있다.
func superAdd(numbers ...int) int {
	total := 0
	for number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}

// 결과 15

// for a := 는 인덱스.
// 그래서 이 함수에선 total에 0부터 5까지 더한다. 함수의 결과는 15.
