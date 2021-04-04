package main

import (
	"fmt"

	"github.com/jhin93/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	// account의 balance만 보이는 func Balance를 적용.
	fmt.Println(account.Balance())
	account.Withdraw(20)
	fmt.Println(account.Balance())
}

// 결과는 다음과 같다.
// 10
// 10
// 아무 변화가 없다.
// 왜냐하면, Withdraw에서 error를 return 했다는 이유로 Go가 관여를 하지 않기 때문. 그렇기에 balance는 여전히 10으로 출력됨.
// 그러나 Go는 자바스크립트의 try catch와 같은 것이 없기에, 개발자가 에러를 알 수 없다.
// 그렇기에 error를 다룰 코드를 따로 작성해주어야만 한다.
