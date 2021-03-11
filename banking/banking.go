package banking

// package banking은 main function을 갖지 않는다. struct만 갖는다. compile(실행)하지 않을 것이기 때문.
// 컴파일은 main.go 가 하고 main.go는 banking.go를 사용한다.

// 그러기 위해선 Export해주어야 하고, 대문자로 작성해야 한다. ex) bankAccount - x, BankAccount- o
// struct 뿐만 아니라, struct의 요소들도 대문자로 해야 export가 가능하다. ex) owner - x, Owner - o
// 대문자 - public. export 가능. 소문자 - private. export 불가능.

// Tip. 뭔가를 export한다면 comment를 해야 한다. comment는 export하는 대상(여기선 Account)으로 시작해야 한다.

// Account struct
type Account struct {
	Owner   string
	Balance int
}
