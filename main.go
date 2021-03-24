package main

import "fmt"

func canIdrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIdrink(20))
}

// 20을 인자로 넣어도 false가 도출된다.
// 10 혹은 18만 리턴되어야 하기 때문.
