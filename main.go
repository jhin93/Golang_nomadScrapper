package main

import "fmt"

// struct. 구조체. ※ Go에는 클래스도, 객체도 없다.
// struct는 자바스크립트의 object와 비슷하면서 map보다 유연한 것이 특징이다.
// 유연한 이유는 밸류의 타입을 map보다 자유롭게 사용 가능하기 때문일 듯. map은 키와 밸류를 각각 지정하면 그대로 사용해야 함.
// struct를 만들기 위해서는 어떤 struct인지 정의해주어야 한다.

func main() {
	type person struct {
		name    string
		age     int
		favFood []string
	}
	food := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: food}
	fmt.Println(nico.favFood)
	// 결과 [kimchi ramen]

	// struct를 작성할 땐 field:value 혹은 value. 둘중 하나만 사용해야 한다. 다음과 같이 섞는 것은 금물
	nico := person{name: "nico", 18, favFood: favFood}

	// field:value 형식으로 작성하는 것을 추천. value만 작성하면 상단의 구조체를 보면서 field를 확인해야 한다.

}
