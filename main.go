package main

import (
	"fmt"

	"github.com/jhin93/learngo/something"
)

func main() {
	fmt.Println("Hello world!") // Go에서는 모듈을 export 하기 위해선 해당 모듈이 대문자로 시작해야 한다. 그래서 Println이 대문자 P로 시작하는 것.
	something.SayHello()        // somethig.go의 SayHello는 대문자로 작성되었기에 export 됨.
	// something.sayBye() // sayBye는 private 함수. 함수가 대문자로 시작하지 않기 때문에 export 되지 않아 사용할 수 없다.
}
