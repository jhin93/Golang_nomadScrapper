package main

import (
	"fmt"

	"github.com/jhin93/learngo/banking"
)

// bankAccount

// 'banking'이라는 package의 Account라는 struct를 사용.
// Account라는 struct를 구조 그대로 사용해야 한다. Account, Owner, Balance 다 대문자로 사용할 것.

func main() {
	account := banking.Account{Owner: "nicolas", Balance: 1000}
	fmt.Println(account)
}
