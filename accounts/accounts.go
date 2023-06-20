package accounts

import (
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a Account) Deposit(amount int) {
	fmt.Println("Gonna deposit", amount)
	a.balance += amount
	// main.go에서 Deposit(10)이 되지 않은 이유 :
	// Deposit 메소드의 a가 main.go의 변수 account의 복사본이기 때문.
	// 그래서 main.go의 account에는 10이 적립되지 않고, accounts.go의 a(main.go의 account 복사본)에 적립된 것.
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}
