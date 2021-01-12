// 1.'main'은 오직 컴파일을 위해 필요한 것. 파일 맨 위가 내가 작성할 패키지의 이름을 작성하는 곳.
package main

import (
	"fmt"

	"github.com/jhin93/learngo/something"
)

// 2. function main이 Go program의 시작점이 되는 부분.
// 컴파일러는 main package와 main function을 먼저 찾고 실행시킨다.
func main() {
	fmt.Println("Hello world")
	// SayHello는 something.go에서 함수가 대문자로 작성되었기에 export가 가능하다.
	// 그렇기에 SayHello는 export되었고 7번째 줄에 import된 모습이다(vscode extension이 자동으로 추가해줌)
	something.SayHello()
	// 하지만 sayBye는 대문자로 시작하지 않기에 private한 함수이다.
	something.sayBye()

}
