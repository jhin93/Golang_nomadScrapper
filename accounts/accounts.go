package accounts

// Account struct
// export를 하기 위해선 주석으로 export하는 대상을 언급해주어야 한다.
// 대문자는 public, 소문자는 private.
type Account struct {
	owner   string
	balance int
}

// NewAccount is account-making function.
// 여기서 NewAccount의 인자를 owner만 받고 balance는 0으로 고정해버리면, 사용자는 새 계좌를 만들 때 임의로 balance를 조작할 수 없다.
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
