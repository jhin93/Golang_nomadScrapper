package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
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
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c) // goroutine으로 부터 받은 메시지를 출력
	}
}

func hitURL(url string, c chan<- result) { // <-c 는 channel로부터 받는다는 뜻인데, chan<-는 channel이 받는 역할만 한다는 뜻이다. 즉 <-c 채널로부터 받는 것은 못하는 것.
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	// 마지막엔 result를 우리 채널로 보낸다.
	c <- result{url: url, status: status}
}
