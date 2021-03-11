package banking

// package banking은 main function을 갖지 않는다. struct만 갖는다. 컴파일하지 않을 것이기 때문.
// 컴파일은 main.go 가 하고 main.go는 banking.go를 사용한다.
// 그러기 위해선 Export해주어야 하고, 대문자로 작성해야 한다. bankAccount - x, BankAccount- o
// Tip. 뭔가를 export한다면 comment를 해야 한다. 주석을 달라는 뜻 (8번째 줄)

// Account struct
type Account struct {
	owner   string
	balance int
}
