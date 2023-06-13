
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
// 0xc0000a6018 0xc0000a6020
```
메모리주소값 출력
```
func main() {
	a := 2
	b := &a
	fmt.Println(*b)
}
// 2
```
변수에 다른값 할당 후 메모리 주소값 출력
```
package main

import "fmt"

func main() {
    a := 2
    b := &a
    a = 5
    fmt.Println(*b)
    // 5
}
```
포인터(*)에 다른 값을 할당해서 원본 데이터값 변경
```
package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 20 // b를 이용해서 a를 업데이트
	fmt.Println(a)
	//20
}
```
slice 사용 예제
```
package main

import "fmt"

func main() {
	//slice는 길이 제한 없는 array
	names := []string{"a", "b", "c"}
	// 요소 추가를 위해선 append 함수 사용
	names = append(names, "ee") // 새로운 slice 를 return하는 것
	fmt.Println(names)
	// [a b c ee]
}

```
*에러케이스 :  
➜  learngo git:(master) ✗ go run main.go  
main.go:6:2: no required module provides package github.com/jhin93/learngo/banking: go.mod file not found in current directory or any parent directory; see 'go help modules'  

해결책 :  
go mod int '소스루트'. 
ex) go mod init github.com/jhin93/learngo. 
