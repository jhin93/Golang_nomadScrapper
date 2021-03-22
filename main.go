package main

import "fmt"

func canIdrink(age int) bool {
	if koreanAge := age + 1; koreanAge > 19 {
		return true
	}
	return false
}

func main() {
	fmt.Println(canIdrink(19))
}
