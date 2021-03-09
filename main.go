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
	// 두번째 방법. 상단의 struct가 안보이거나 해도 어떤 식으로 이루어져있는지 알 수 있다.
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	//fmt.Println(nico) 혹은 fmt.Println(nico.age)
	fmt.Println(nico.age)
}
