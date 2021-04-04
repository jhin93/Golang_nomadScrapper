package accounts

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
