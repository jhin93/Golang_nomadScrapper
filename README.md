
[Golang 스탠다드 라이브러리 패키지]  
https://pkg.go.dev/std  

함수(func) 여러 return value 반환 에시

```golang
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

```golang
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
```golang
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
```golang
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
```golang
func main() {
	a := 2
	b := &a
	fmt.Println(*b)
}
// 2
```
변수에 다른값 할당 후 메모리 주소값 출력
```golang
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
```golang
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
```golang
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


포인터(pointer(*)) 타입은 '주소를 저장'하는 타입이다.  
```golang
// NewAccount creates Account
func NewAccount(owner string) *Account { 
	// 주소를 저장하는 타입을 'pointer 타입'이라고 한다. 'pointer 타입' 은 '*'과 주소를 가진 데이터의 형태가 결합된 형태다(ex var ptr *int).
	// 이 함수의 리턴 타입인 *Account는 '구조체 Account 형태의 데이터의 주소'를 의미한다.
	// 그렇기 때문에 변수 account(구조체 Account 형태의 데이터)의 주소, 즉 &account가 리턴값이 되어야 한다.
	account := Account{owner: owner, balance: 0}
	return &account
}

```
포인터(pointer) 타입 예시  
```golang
var x int = 42
var ptr *int = &x

fmt.Println(*ptr) // 출력: 42

fmt.Println(ptr)  // 출력: 0xc00010e008(x 변수의 주소). ptr은 x 변수의 주소를 저장
fmt.Println(&ptr) // 출력: 0xc000106018(ptr 변수의 주소). &ptr은 ptr 변수 자체의 주소
```

메소드(Method)
```golang
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

메소드 2(직접 만든 type에 메소드를 적용하는 예제)
```golang
mydict.go
	package mydict
	
	import "errors"
	
	// Dictionary type
	type Dictionary map[string]string
	
	var errNotFound = errors.New("Not Found")
	
	// Search for a word
	func (d Dictionary) Search(word string) (string, error) {
		value, exists := d[word] // value는 찾고자 하는 값(d[word])이고 exists는 boolean이다.
		if exists {
			// 찾는 value가 있다면, error 값은 nil로 반환한다.
			//리턴 형태는 (string, error) 이고, 찾는 값이 존재해서 에러가 없으니까 nil로 반환하는 것.
			return value, nil
		}
		return "", errNotFound // 찾는 값이 없으니까 string은 "", 에러는 미리 만들어둔 에러변수를 반환한다.
	}

main.go
	package main

	import (
		"fmt"
	
		"github.com/jhin93/learngo/mydict"
	)
	
	func main() {
		dictionary := mydict.Dictionary{"first": "First word"}
		definition, err := dictionary.Search("second")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition)
		}
	}

//결과 : Not Found



```

조건문 
```golang
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
```golang
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
// 'String'메소드는 fmt.Println()을 호출할 때 Go가 내부적으로 자동으로 호출하는 메소드이다. 이를 위와 같이 변형하면, fmt.Println() 실행시 적용된다.
```


고루틴(Goroutine) 예시
```golang
package main

import (
	"fmt"
	"time"
)

func calculateSquare(number int) {
	result := number * number
	fmt.Printf("Square of %d is %d\n", number, result)
}

func calculateCube(number int) {
	result := number * number * number
	fmt.Printf("Cube of %d is %d\n", number, result)
}

func main() {
	number := 5

	// calculateSquare 함수를 goroutine으로 실행
	go calculateSquare(number)

	// calculateCube 함수를 goroutine으로 실행
	go calculateCube(number)

	// 메인 함수의 실행이 종료되지 않도록 대기
	time.Sleep(1 * time.Second)
}

```
이 예시에서는 calculateSquare 함수와 calculateCube 함수를 두 개의 goroutine으로 실행합니다. 각 함수는 간단한 계산을 수행하여 결과를 출력합니다. 메인 함수에서는 두 함수를 goroutine으로 실행한 후, time.Sleep을 사용하여 메인 함수의 실행이 종료되지 않도록 합니다. 결과적으로, Square과 Cube 함수가 동시에 실행되어 5의 제곱과 세제곱이 출력됩니다. 출력 결과는 순서가 일정하지 않을 수 있으며, 이는 goroutine 스케줄링에 따라 달라질 수 있습니다.  
```golang
Square of 5 is 25
Cube of 5 is 125
또는
Cube of 5 is 125
Square of 5 is 25
```

채널(Channel)  
채널은 주로 <- 연산자를 사용하여 데이터를 보내고 받습니다.  
데이터를 채널에 보내기 위해서는 채널 <- 데이터와 같이 사용하고(ex c <- result),   
채널에서 데이터를 받기 위해서는 데이터 <- 채널과 같이 사용합니다(ex result <- c).  
```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	// channel은 goroutine과 메인함수, 혹은 goroutine 간의 커뮤니케이션을 하는 방법이다.
	// channel(chan)을 isSexy로 보내고, 그로 인해 isSexy는 메인함수랑 커뮤니케이션 할 수 있게 된다.
	// 채널 만드는 법
	// 변수 := make(chan 채널을통해보낼타입)
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	// 채널을 두 함수(isSexy nico, isSexy flynn)로 보냄.
	for _, person := range people {
		go isSexy(person, c)
	}
	// time.Sleep()이 없어도 메인함수가 바로 종료되지 않고 기다린다.
	// 채널로부터 뭔가를 받을 때, 메인함수는 뭔가 답을 받을때까지 기다린다.
	fmt.Println(<-c)
	fmt.Println(<-c)
	// fmt.Println(<-c) 이런식으로 goroutine을 초과해서 실행하려 하면 deadlock이 뜬다. 이미 모든 goroutine이 끝났기 때문.
}

func isSexy(person string, c chan bool) { // 채널을 통해 보낼 타입이 bool이기에 같이 적어줌(c chan bool).
	time.Sleep(time.Second * 5) // 두 함수(isSexy nico, isSexy flynn)는 5초 후에 true 라는 메시지를 나에게 2개 보낸다.
	fmt.Println(person)
	c <- true // goroutine으로부터 return을 받는 대신에 채널을 통해 메시지를 전달한다.
}

// 결과 :
// nico
// true
// flynn
// true
// 둘은 엄밀히 동시에 끝난 것.

```

채널 루프
```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c) // 5개의 메세지 리시버를 만드는 것.
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- person + " is sexy"
}

// 결과
// dal is sexy
// flynn is sexy
// nico is sexy
// japanguy is sexy
// larry is sexy

```

goroutine + 채널 + type + range + 반복문(loop)  
```golang
package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) { // c(변수 이름) chan<-(메시지 받기만 하는 채널) result(메시지 형태)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	// 마지막엔 result를 우리 채널로 보낸다.
	c <- requestResult{url: url, status: status}
}

```
goquery  
https://github.com/PuerkitoBio/goquery  
go가 html 내부를 들여다볼 수 있도록 해주는 라이브러리.  


doc.Find() 예시 - 아래 링크에서 'doc.Find'로 검색했을 때 2번째 결과  
https://pkg.go.dev/github.com/PuerkitoBio/goquery#readme-examples  


fmt.Sprintf  
fmt.Sprintf 함수는 형식화된 문자열을 생성하는 기능을 제공하는 Go 언어의 함수입니다.  
일반적으로, fmt.Sprintf 함수는 형식 문자열과 해당 형식 문자열에 대응하는 값들을 인자로 받습니다.  
그런 다음, 형식 문자열에서 % 기호와 특정 문자를 사용하여 값을 대체하고, 새로운 문자열을 반환합니다.  
예를 들어, 다음은 fmt.Sprintf 함수를 사용하여 형식화된 문자열을 생성하는 예시입니다:  
```golang
name := "Alice"
age := 30
formattedString := fmt.Sprintf("My name is %s and I'm %d years old.", name, age)
fmt.Println(formattedString)

// 위의 예시에서 %s는 문자열을 대체하기 위한 형식 문자열입니다. %d는 10진수 정수를 대체하기 위한 형식 문자열입니다.
// fmt.Sprintf 함수는 name 변수를 %s에 대응하여 문자열로 대체하고, age 변수를 %d에 대응하여 정수로 대체합니다.
// 그런 다음, formattedString 변수에는 "My name is Alice and I'm 30 years old."라는 형식화된 문자열이 저장됩니다. 

```

strings.Replace  

일반적으로, strings.Replace 함수는 세 가지 인자를 받습니다: 1. 원본 문자열 2. 대체할 문자열 3. 대체 횟수입니다.  
이 함수는 원본 문자열에서 대체할 문자열을 찾아 해당 위치에 새로운 문자열을 삽입하여 새로운 문자열을 생성합니다.  

```golang
str := "Hello, World!"
newStr := strings.Replace(str, "Hello", "Hi", 1)
fmt.Println(newStr)
//  "Hi, World!"
```

```golang
str := "apple apple apple"
newStr := strings.Replace(str, "apple", "orange", 2)
fmt.Println(newStr)

// 위의 예시에서 strings.Replace 함수는 str 문자열에서 "apple"을 찾아 "orange"로 대체합니다.
// 대체 횟수를 2로 설정했으므로, 처음 두 번의 "apple"만 대체됩니다. 따라서 결과는 "orange orange apple"가 출력됩니다.
// 마지막 "apple"은 대체 횟수에 도달하지 않아 그대로 남게 됩니다.
```

echo로 서버 만들기  
https://echo.labstack.com/docs/quick-start#hello-world  

