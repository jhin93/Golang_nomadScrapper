package main

import (
	"errors"
	"fmt"
)

type result struct {
	url    string
	status int
}

var errRequestFailed = errors.New("Request failed")

func main() {
	results := make(map[string]string)
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		hitURL(url, c)
	}
}

// 이 채널은 데이터를 받을 순 없고 보낼 수만 있다.
func hitURL(url string, c chan<- result) { // send-only를 'chan<-' 로 표현한다.
	fmt.Println("Checking:", url)
	fmt.Println(<-c)
}
