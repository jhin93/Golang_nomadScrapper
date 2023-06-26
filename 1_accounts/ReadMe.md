
함수(func) 여러 return value 반환 에시

```java
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

```java
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
```java
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
```java
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
```java
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
go mod init '소스루트' 
ex) go mod init github.com/jhin93/learngo 


pointer(*) 타입은 '주소를 저장'하는 타입이다.  
```java
// NewAccount creates Account
func NewAccount(owner string) *Account { 
	// 주소를 저장하는 타입을 'pointer 타입'이라고 한다. 'pointer 타입' 은 '*'과 주소를 가진 데이터의 형태가 결합된 형태다(ex var ptr *int).
	// 이 함수의 리턴 타입인 *Account는 '구조체 Account 형태의 데이터의 주소'를 의미한다.
	// 그렇기 때문에 변수 account(구조체 Account 형태의 데이터)의 주소, 즉 &account가 리턴값이 되어야 한다.
	account := Account{owner: owner, balance: 0}
	return &account
}

```
pointer 타입 예시  
```java
var x int = 42
var ptr *int = &x

fmt.Println(*ptr) // 출력: 42

fmt.Println(ptr)  // 출력: 0xc00010e008 ptr은 x 변수의 주소를 저장
fmt.Println(&ptr) // 출력: 0xc000106018 &ptr은 ptr 변수 자체의 주소
```

메소드(Method)
```java
// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// '메소드'는 하나를 제외하고 함수와 작성방식이 동일하다
// func와 함수이름 사이에 Account 타입의 변수 'a'와 그 타입 'Account'를 적어준다. 이때 'a'를 'receiver' 라고 한다.
// receiver 작성 규칙 1. struct의 첫글자를 따서 소문자로 지어야 한다.
// Deposit x amount on your account
func (a Account) Deposit(amount int) {
	a.balance += amount
}
```

조건문 
```java
import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // 이런식으로 조건문 안에 변수를 만드는 게 오직 조건문에서만 사용되는 변수라는 걸 알려준다.
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
```

'func'와 'method'의 차이점  
func은 독립적인 함수를 정의하는 데 사용되고, method는 특정 타입과 관련된 함수를 정의하는 데 사용됩니다. func은 어떤 타입에도 속하지 않고 독립적으로 사용될 수 있으며, method는 특정 타입과 연결되어 해당 타입의 변수에서 호출됩니다.


Go가 내부적으로 호출하는(String)을 메소드를 사용하는 방법  
```java
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
// 'String'메소드는 fmt.Println()을 호출할 때 Go가 내부적으로 자동으로 호출하는 메소드이다. 이를 위와 같이 변형하면, fmt.Println() 실행시 적용된다.
```
