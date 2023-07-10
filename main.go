package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

// var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"
// var baseURL string = "https://www.airbnb.com/"
// var baseURL string = "https://dict.naver.com/"

func main() {
	// 1. 페이지들을 받아오기
	// 2. 각각의 페이지 방문하기
	// 3. 페이지로부터 job을 추출하기
	// 4. 추출한 job을 엑셀에 집어넣기
	getPages()

}

// 얼마나 많은 페이지가 있는지 알려준다.
func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	// 이 getPages 함수가 끝났을 때 아래 goquery 문서를 닫아야 함.
	// 왜냐하면 goquery의 res.Body는 기본적으로 IO(input & output)이기 때문.
	// 이렇게 메모리가 새어나가는 걸 막을 수 있음.
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	fmt.Println(doc)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

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
