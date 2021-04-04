package accounts

import "errors"

// Account struct
// export를 하기 위해선 주석으로 export하는 대상을 언급해주어야 한다.
// 대문자는 public, 소문자는 private.
type Account struct {
	owner   string
	balance int
}

// 코드 퀄리티를 위해 에러변수를 작성할때는 'err땡땡땡' 이어야 한다.
var errNoMoney = errors.New("Can't withdraw")

// NewAccount is account-making function.
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	// 위와 같은 작성 방식을 'pointer receiver'라고 함.
	// 여기서 포인터(*)를 사용하면 (a *Account) Deposit 메서드를 호출한 account를 사용하라는 뜻.
	a.balance += amount
}

// Balance of your account.
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	// 잔액이 마이너스 일 경우 출금을 방지하는 로직.
	if a.balance < amount {
		// error를 리턴한다고 명시해야 함. ...Withdraw(amount int) 'error'...
		return errNoMoney
	}
	a.balance -= amount
	// 조건문 밖에서도 리턴을 해주어야 한다.
	// boolean에 true, false 처럼 error에도 두가지 value가 있다. 바로 error와 nil. nil은 자바스크립트의 null과 유사하다.
	return nil
}
