package main

import (
	"fmt"
	"time"
)

func main() {
	// channel은 goroutine과 메인함수, 혹은 goroutine 간의 커뮤니케이션을 하는 방법이다.
	// channel(chan)을 isSexy로 보내고, 그로 인해 isSexy는 메인함수랑 커뮤니케이션 할 수 있게 된다.
	// 채널 만드는 법
	// 변수 := make(chan 채널을통해보낼타입)
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people { // 채널을 두 함수(isSexy nico, isSexy flynn)로 보냄.
		go isSexy(person, c)
	}
	result := <-c // goroutine으로부터 보내진 메시지를 받는 방법
	fmt.Println(result)
}

func isSexy(person string, c chan bool) { // 채널을 통해 보낼 타입이 bool이기에 같이 적어줌(c chan bool).
	time.Sleep(time.Second * 5) // 두 함수(isSexy nico, isSexy flynn)는 5초 후에 true 라는 메시지를 나에게 2개 보낸다.
	c <- true                   // goroutine으로부터 return을 받는 대신에 채널을 통해 메시지를 전달한다.
}
