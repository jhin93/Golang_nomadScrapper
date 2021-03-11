package main

import (
	"fmt"

	"github.com/jhin93/learngo/accounts"
)

// bankAccount

// 다른 사람들이 Account를 만들지 못하게 할 것. Balance를 조정할 수 없도록.
// account를 생성하는 function을 만들 것. Go의 흔한 패턴.
// Go에는 constructor가 없기 때문에 function으로 construct하거나 struct를 만드는 것.

// accounts package의 NewAccount를 import해온다.
func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
	// 결과: &{nico 0}. nico는 &이 붙은 address다. nico가 복사본이 아니라
	// 만약 func NewAccount에서 *과 &를 사용하지 않으면 결과는 {nico 0}.
	// account.balance = 100000 - 이제 이런 것이 불가하다. 왜냐하면 Account struct에서 balance는 private하기 떄문. owner도 마찬가지
}

// Go에서 constructor를 만드는 방법.
// function을 만들고(func NewAccount), object를 리턴(account 메모리의 값. *Account)
