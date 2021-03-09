package main

import "fmt"

// struct. go는 class도 객체도 없다.
// struct는 자바스크립트의 object와 비슷하면서 map보다 유연한 것이 특징이다.
// struct를 만들기 위해서는 어떤 struct인지 정의해주어야 한다.
type person struct {
	name    string
	age     int
	favFood []string // slice
}

func main() {
	// struct를 사용할 땐, field:value 혹은 value. 둘 중 하나로 통일해서 작성해야만 한다.
	// 다음과 같이 field:value와 value를 섞어서 사용하면 안된다. nico := person{name: "nico", 18, favFood: favFood}
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	//fmt.Println(nico) 혹은 fmt.Println(nico.age)
	fmt.Println(nico.age)
}
