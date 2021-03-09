package main

import "fmt"

// struct. 메서드 역시 struct에 포함될 수 있다.
// Go에는 constructor method가 없다. 스스로 constructor를 실행해야 한다.
// 자바스크립트 constructor. https://developer.mozilla.org/ko/docs/Web/JavaScript/Reference/Classes/constructor
type person struct {
	name    string
	age     int
	favFood []string // slice
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico.age)
}
