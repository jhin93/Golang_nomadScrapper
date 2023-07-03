package main

import (
	"fmt"
	"time"
)

func main() {
	// Goroutines이란 기본적으로 다른 함수와 동시에 실행시키는 함수
	// 아래처럼 go가 둘다 붙어있으면, 첫번째 Goroutines과 두번째 Goroutines이 실행되고 끝나버린다. 더 진행할 작업이 없기 때문이다.
	// Goroutines은 프로그램이 실행되는 동안, 즉 메인함수가 실행되는 동안만 유효하다.
	// 메인함수가 끝나버리면 goroutines도 소멸된다. 다 go가 붙어있으면, 메인함수는 할 게 없어진다.
	go sexyCount("nico")
	go sexyCount("flynn")
	time.Sleep(time.Second * 5) // 이 time.Sleep()함수가 5초간 실행되는 동안 위 goroutine들은 timse.Sleep()이 실행되는 동안만 같이 실행된다.
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
