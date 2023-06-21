package main

import (
	"fmt"

	"github.com/jhin93/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err) // Fatalln은 Println을 호출하고 프로그램을 종료시키는 역할
	} // 결과 : Can't withdraw you are poor
	fmt.Println(account.Balance())
}
