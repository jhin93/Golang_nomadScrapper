package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println("Waiting for messages")
	resultOne := <-c                                   // 첫번째 메세지를 기다림. 그게 뭐든 간에 resultOne에 넣음
	resultTwo := <-c                                   // 두번째 메세지를 기다림. 그게 뭐든 간에 resultOne에 넣음
	fmt.Println("Received this message : ", resultOne) // <-c는 채널로부터 메시지를 얻고 있다는 뜻. 메세지 한개를 얻으면, 다음 줄로 넘어감.
	// <-c 는 메세지를 기다리는 것. 이걸 blocking operation이라고 부른다. 프로그램이 메세지를 받을 때까지 멈추는 것.
	fmt.Println("Received this message : ", resultTwo) // 이 줄의 코드도 메세지를 기다리라는 뜻. 메세지를 받을 때까지 기다리고 받으면 그 다음 줄을 진행함.
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- person + " is sexy"
}
