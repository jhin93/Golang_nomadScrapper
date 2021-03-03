package main

import "fmt"

// 데이터 구조 : map
// 자바스크립트의 object와 비슷하다.

func main() {
	// map도 range를 이용해서 반복문에 이용할 수 있다.
	nico := map[string]string{"name": "nico", "age": "12"}
	// key나 value 둘중 하나를 ignore 시킬 수도 있다.
	for key, _ := range nico {
		fmt.Println(key)
	}
}
