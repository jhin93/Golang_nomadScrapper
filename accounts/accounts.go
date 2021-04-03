package accounts

import "fmt"

// Account struct
// export를 하기 위해선 주석으로 export하는 대상을 언급해주어야 한다.
// 대문자는 public, 소문자는 private.
type Account struct {
	owner   string
	balance int
}

// NewAccount is account-making function.
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// 메소드 작성하기.
// func와 함수명 사이에 receiver를 적어준다.
// Go에서는 receiver가 메소드와 동등하다.

// receiver를 작성 규칙.
// 1. function의 이름 전에 작성한다.
// 2. struct의 첫 글자를 따서 소문자로 지어야 한다.
// ex) func (a Account) Deposit. --> receiver의 이름은 a이고 a의 타입은 Account라는 struct이다.

// Deposit x amount on your account
// 이 Deposit은 리턴할 값을 쓰지 않았다.
func (a Account) Deposit(amount int) {
	// 메서드 내에서 a로 접근을 하면 owner, balance를 변경할 수 있다.
	fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

// Balance of your account. account의 balance만 보이도록 method를 작성.
func (a Account) Balance() int {
	return a.balance
}
