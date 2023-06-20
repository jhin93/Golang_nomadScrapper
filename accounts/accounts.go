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
	// 주소를 저장하는 타입을 'pointer 타입'이라고 한다. 'pointer 타입' 은 '*'과 주소를 가진 데이터의 형태가 결합된 형태다(ex var ptr *int).
	// 이 함수의 리턴 타입인 *Account는 '구조체 Account 형태의 데이터의 주소'를 의미한다.
	// 그렇기 때문에 변수 account(구조체 Account 형태의 데이터)의 주소, 즉 &account가 리턴값이 되어야 한다.
	account := Account{owner: owner, balance: 0}
	return &account
}

// '메소드'는 하나를 제외하고 함수와 작성방식이 동일하다
// func와 함수이름 사이에 Account 타입의 변수 'a'와 그 타입 'Account'를 적어준다. 이때 'a'를 'receiver' 라고 한다.
// receiver 작성 규칙 1. struct의 첫글자를 따서 소문자로 지어야 한다.
// Deposit x amount on your account
func (a Account) Deposit(amount int) {
	fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}
