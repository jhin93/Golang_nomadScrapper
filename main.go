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
	// account.Withdraw()로부터 error를 받아온다.
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}

// 실행결과
// 10
// Can't withdraw you are pooor
// 10

// err := .... -> 앞으로 흔하게 쓰게 될 것. 항상 error를 체크해야 한다.
