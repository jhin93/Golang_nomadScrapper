package main

import "fmt"

func canIdrink(age int) bool {
	if koreanAge := age - 1; koreanAge >= 19 {
		return true
	}
	return false
}

func main() {
	fmt.Println(canIdrink(19))
}

// 20을 인자로 넣어도 false가 도출된다.
// 10 혹은 18만 리턴되어야 하기 때문.
