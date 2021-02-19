package main

import (
	"fmt"
	"strings"
)

// naked return
// 함수를 작성할 때, 마지막에 return할 variable을 꼭 명시하지 않아도 된다. ex return 어쩌구저쩌구..
// 함수 작성할 시 {} 이전에 return할 값을 넣는 곳에 처음부터 쓰고 시작해도 된다. 대신 리턴할 variable과 함수 내부에 작성될 variable은 같아야 한다.
// 대신 마지막에 'return'은 꼭 작성해준다.

func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, upper := lenAndUpper("nico")
	fmt.Println(totalLength, upper)
}
