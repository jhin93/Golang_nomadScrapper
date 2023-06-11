package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLenth, _ := lenAndUpper("nico") // _는 value를 무시한다.
	fmt.Println(totalLenth)
}

// result : 4
