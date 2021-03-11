package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

// struct도 private이 될 수 있고(Account -> account),
// public인 function을 만들 수 있다. 이 function이 struct를 만들 것.

// NewAccount creates Account. NewAccount는 owner만을 갖는다. 그리고 account의 복사본을 리턴한다. *를 붙여서. *는 살펴보는 기능이다.
func NewAccount(owner string) *Account {
	// 1. account를 초기화한다.
	account := Account{owner: owner, balance: 0}
	// 2. 복사본의 address(&)를 return한다. 새로운 object를 리턴하고 싶은데 만들긴 싫으니 메모리 주소를 return. 항상 값 복사를 하면 데이터가 늘어나 느려진다.
	return &account
	// 3. &account = account의 메모리. 'acount의 메모리'의 '값'을 *을 붙인 Account를 리턴함으로써 알 수 있다.
	// *를 사용하면 메모리 주소의 값 만을 살펴보는 것이라서, 메모리 주소에 새로운 값이 대입되도 그것을 확인할 수 있다.
	// 추가로,단순히 값 복사를 하면 복사하기 바로 전 시점의 메모리 값을 복사하기에 그 이후에 새로 메모리 값이 변해도 그것을 알 수 없다.
}
