
함수(func) 여러 return value 반환 에시

```
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLenth, upperName := lenAndUpper("nico")
	fmt.Println(totalLenth, upperName)
}

// result : 4 NICO
```

'defer'는 함수 실행이 종료된 이후 필요한 로직을 실행하게 해주는 기능이다.

```
import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	// defer는 함수 실행이 종료되고 나서 진행되는 로직임.
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLenth, uppercase := lenAndUpper("nico") // _는 value를 무시한다.
	fmt.Println(totalLenth, uppercase)
}

// result
// I'm done - 함수 lenAndUpper 가 실행되고 난 이후 defer에 의해 출력됨
// 4 NICO - 함수 main에 의해 출력
```
포인터
```
package main

import "fmt"

func main() {
	a := 2
	b := a // b는 a의 value를 복사한다.
	a = 10
	fmt.Println(a, b)
	// result 10 2
}
```
메모리주소
```
package main

import "fmt"

func main() {
	a := 2
	b := 5
	fmt.Println(&a, &b)
}
```
