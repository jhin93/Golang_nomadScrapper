package main

import (
	"fmt"
)

// 변수(Variables)와 상수(Constants)
// 변수는 말 그대로 변수, 값을 변경할 수 있다. 하지만 상수는 변수지만 값을 바꿀 수 없다.
// Variables는 let, Constants는 const로 보면 된다.

func main() {
	//const name = "nico"
	//	이것은 타입이 없는 상수(const name untyped string = "nico").
	//	Go는 type 언어다. 타입이 무엇인지 알려줘야 함. String인지 Boolean인지
	//const name string = "nico"
	//	이제 name이 string타입이란 걸 알 수 있다.
	//var name string = "nico"
	//name = "lynn"
	//	var는 변수이기에 재할당이 가능하다.
	//name := "nico"
	//	var name string = "nico"를 위 코드처럼 축약시킨다면 Go가 type을 찾아줌. 정해진 type은 임의로 변경 x
	name := "nico"
	name = "lynn"
	//	만약 name의 type이 위와 같이 Boolean(false)로 정해진다면 그 name 변수는 String으로 변환할 수 없다는 메시지가 뜸.
	//	왜냐하면 첫번째 변수의 type에 의존해서 두번째 변수의 type도 정해진다.
	//	또한 func을 벗어나면 축약형 코드(ex name := false)는 작동하지 않는다.
	//	그리고 축약형은 변수에만 적용 가능하다. 축약형이 존재하면 Go가 첫번째 변수를 기준으로 변수의 type을 찾아 정해준다.
	//	함수 밖에서 코드가 작동하게 하려면 var name bool = false <- 이런식으로 전부 작성해야 작동한다.
	fmt.Println(name)
}
