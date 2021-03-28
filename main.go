package main

import (
	"fmt"

	"github.com/jhin93/learngo/accounts"
)

// GO에서 constructor를 만드는 방법

// 1. function(ex NewAccount)을 만들어서 object를 return한다. 복사본이 아닌 실제 메모리 address를 return. 항상 복사본을 만드는 것은 지양.
//
func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
	// 결과
	// &{nico 278}
	account.balance = 1000
	// 위의 문장은 불가하다. 왜냐하면 import되는 NewAccount에서는 오직 owner만 인자로 받기 때문.
}
