package main

import (
	"fmt"
	"time"
)

func main() {
	// 메인함수는 내부의 각종 함수들의 결과를 저장하는 곳. goroutines와 커뮤니케이션하는 방법이 있다.
	go sexyCount("nico")
	go sexyCount("flynn")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
