package main

import (
	"log"
	"net/http"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	// 1. 페이지들을 받아오기
	// 2. 각각의 페이지 방문하기
	// 3. 페이지로부터 job을 추출하기
	// 4. 추출한 job을 엑셀에 집어넣기
	pages := getPages()

}

// 얼마나 많은 페이지가 있는지 알려준다.
func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	return 0
}

// 에러 체크 함수
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// res 이상 체크 함수
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
