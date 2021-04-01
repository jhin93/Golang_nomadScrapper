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
