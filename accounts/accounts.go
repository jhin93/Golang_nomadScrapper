package accounts

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
