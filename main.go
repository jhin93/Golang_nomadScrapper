package main

import "fmt"

// 데이터 구조 : map
// 자바스크립트의 object와 비슷하다.

func main() {
	// 작성하는 법
	// map[key의 타입]value의 타입{key:value}
	// 아래 map의 key는 string, value도 string.
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)
}
