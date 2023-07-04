package main

import (
	"time"
)

func main() {
	// channel은 goroutine과 메인함수, 혹은 goroutine 간의 커뮤니케이션을 하는 방법이다.
	// channel(chan)을 isSexy로 보내고, 그로 인해 isSexy는 메인함수랑 커뮤니케이션 할 수 있게 된다.
	// 채널 만드는 법
	// 변수 := make(chan 채널을통해보낼타입)
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c) // 이런식으로 채널(c)을 isSexy로 보내줌
	}
}

func isSexy(person string, c chan bool) { // 채널을 통해 보낼 타입이 bool이기에 같이 적어줌(c chan bool).
	time.Sleep(time.Second * 5)
	c <- true // c라는 channel에다가 true를 보내준다.
}
