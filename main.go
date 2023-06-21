package main

import (
	"fmt"

	"github.com/jhin93/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	account.Withdraw(20) // 에러를 리턴했다는 이유로 Go가 관여를 하지 않음
	fmt.Println(account.Balance())
}
