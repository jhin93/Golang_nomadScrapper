package main

import "fmt"

// struct.
// struct는 자바스크립트의 object와 비슷하면서 map보다 유연한 것이 특징이다.
// struct를 만들기 위해서는 어떤 struct인지 정의해주어야 한다.
type person struct {
	name    string
	age     int
	favFood []string // slice
}

func main() {
	// person에게 value를 주는 방법에는 두가지가 있다.
	// 첫번째로 순서대로 작성하는 방법. name -> age -> favFood
	// 그러나 이렇게 작성하는 것은 그다지 추천하지 않는다. 왜냐하면 상단에 struct를 봐야 어떤 field가 무슨 value인지 알 수 있다.
	favFood := []string{"kimchi", "ramen"}
	nico := person{"nico", 18, favFood}
	//fmt.Println(nico) 혹은 fmt.Println(nico.age)
	fmt.Println(nico.age)
}
