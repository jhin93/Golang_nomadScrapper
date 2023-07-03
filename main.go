package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	// 패닉(panic) 예시
	// 패닉은 컴파일러가 잡아내지 못하는 에러이다.
	// panic: assignment to entry in nil map
	// 초기화 되지 않은 map에 특정 값을 할당할 수 없기 때문에 일어난 에러이다. 초기화되지 않은 map은 nil이다.
	// 초기화를 해줘야 빈 map이 되고 어떤 값이든 넣을 수 있게 된다.
	// 초기화 예시 1. var results = map[string]string{} <- 끝부분에 {} 를 붙인다.
	// 초기화 예시 2. var results = make(map[string]string) <- make() 함수 안에 넣는다.
	var results map[string]string
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
	results["hello"] = "Hello"
	for _, url := range urls {
		hitURL(url)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	// 에러가 nil이
	// response가 300대까지는 리다이렉션을 해주고, 400부터는 뭔가 문제가 있다는 뜻
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
