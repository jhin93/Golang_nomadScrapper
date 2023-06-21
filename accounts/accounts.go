package accounts

import (
	"errors"
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
func (a *Account) Deposit(amount int) {
	fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error { // return 타입이 error 라고 명시해줘야 조건문 안의 return이 error가 된다.
	if a.balance < amount {
		return errors.New("Can't withdraw you are poor")
	}
	a.balance -= amount
	return nil // nill은 자바스크립트의 null과 같다
}
