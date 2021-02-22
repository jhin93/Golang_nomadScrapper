package main

import "fmt"

// Go에서 loop는 오로지 for를 사용하는 것으로만 가능하다.
// range는 array에 loop를 적용할 수 있도록 해준다.
func superAdd(numbers ...int) int {
	for number := range numbers {
		fmt.Println(number)
	}
	return 2 // return type이 int이라서 return 1 을 하는 것.
}

func main() {
	superAdd(1, 2, 3, 4, 5, 6)
}
