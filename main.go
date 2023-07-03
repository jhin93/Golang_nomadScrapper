package main

import (
	"fmt"
	"time"
)

func main() {
	// Goroutines이란 기본적으로 다른 함수와 동시에 실행시키는 함수
	go sexyCount("nico") // go 를 붙임으로써 sexyCount("nico")를 마치고 sexyCount("flynn")가 실행되던 것에서 두 함수 동시 실행으로 바뀌었다.
	sexyCount("flynn")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
