// 1.'main'은 오직 컴파일을 위해 필요한 것. 파일 맨 위가 내가 작성할 패키지의 이름을 작성하는 곳.
package main

import (
	"fmt"
)

// 2. function main이 Go program의 시작점이 되는 부분.
// 컴파일러는 main package와 main function을 먼저 찾고 실행시킨다.
func main() {
	fmt.Println("Hello world!")
}
