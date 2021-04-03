package main

import (
	"fmt"

	"github.com/jhin93/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	// account의 balance만 보이는 func Balance를 적용.
	fmt.Println(account.Balance())
}

// 위의 결과는 0이 나온다. 그 이유는 다음과 같다.
// Go에서 object와 struct에 관여하는 부분 때문이다. 바로 복사본을 만드는 것.

// 1. account.Deposit()을 하면, Go는 account를 받아온다.
// 2. accounts.go 파일의 func (a Account) Deposit 에서, (a Account) --> 여기서 복사본을 만들어버린다.
// 3. Deposit에서 사용되는 a는 account가 맞지만, main.go 파일의 account.Deposit --> 여기의 account의 복사본이다.
// 실제 account가 아니라는 말.

// 이는 method를 보호할 때 유용한다. ex) method에 속해있는 object가 변경되는 결 원치 않을 때
// 예를 들어 accounts.go 파일의 Balance() 메소드는 a가 account의 복사본이어도 상관하지 않는다. 단지 balance를 불러오기만 하기 때문에.

// 그러나 Deposit 메소드의 경우는 다르다.
// 복사본을 만들지 않고 account의 balance를 증가시키는 것. 복사본을 만들어버리면 실제로 반영되지 않고 위처럼 0이 도출된다.
// 이를 해결하기 위해선 Deposit 메소드의 receiver의 타입에 *를 추가하는 것이다. ex) a *Account

// 위와 같이 포인터를 적용하면 account.Deposit()를 호출할 때, account를 복사하지 않고 Deposit method를 호출한 account를 사용하는 것.
