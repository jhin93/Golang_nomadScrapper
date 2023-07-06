package main

import (
	"errors"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
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
		result := <-c
		results[result.url] = result.status
		// goroutine으로 부터 받은 메시지를 출력
		// 이 반복문이 끝나는 시간은, url 체크하는데 가장 오래걸리는 사이트 체크하는 시간과 동일. 병렬적으로 실행되는 것.
		// 가장 늦은 메세지를 받으면 프로그램이 완료되는 것.
	}
}

func hitURL(url string, c chan<- requestResult) { // c(변수 이름) chan<-(메시지 받기만 하는 채널) result(메시지 형태)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	// 마지막엔 result를 우리 채널로 보낸다.
	c <- requestResult{url: url, status: status}
}
