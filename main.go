package main

import (
	"fmt"
	"strings"
)

// defer
// defer는 function이 값을 return하고 나면 실행된다.
// function이 실행되고 나서 코드가 실행된다.
// 아래 코드는 다음과 같은 순서로 코드가 작동한다.
// 1. lenAndUpper 함수가 실행된다. 그리고 return이 되지만 fmt.Println이 없어 티가 안남.
// 2. lenAndUpper의 defer가 실행된다. lenAndUpper가 끝났으니까.
// 3. func main의 fmt.Println가 실행된다.
// 4. func main의 defer 가 실행된다. 이 defer가 속한 'main'함수 실행이 끝났으니까.

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	defer fmt.Println("I'm last")
	totalLength, upper := lenAndUpper("nico")
	fmt.Println(totalLength, upper)
}
