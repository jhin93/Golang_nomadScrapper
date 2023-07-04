package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c) // 5개의 메세지 리시버를 만드는 것.
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- person + " is sexy"
}
