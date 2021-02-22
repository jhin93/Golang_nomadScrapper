package main

import (
	"fmt"
	"strings"
)

// defer
// defer는 function이 값을 return하고 나면 실행된다.
// function이 실행되고 나서 코드가 실행된다.
// 아래 코드에서는 다음과 같은 순서로 코드가 작동한다.
// 1. 17번째 줄의 lenAndUpper 함수가 실행된다. 실행은 되었지만 fmt.Println이 함수 내부에 없어 티가 나지 않는다.
// 2. 18번째 줄의 defer 가 실행된다.
// 3. 24번째 줄의 func main()가 실행된다. 이때 fmt.Println(totalLength, upper)가 실행된다.
// 4. 25번째 줄의 defer 가 실행된다. 이 defer가 속한 'main'함수가 끝났으니까.

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
