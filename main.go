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
