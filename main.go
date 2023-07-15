package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=python&recruitPage=1&recruitSort=relation&recruitPageCount=40&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10&show_applied=&quick_apply=&except_read=&ai_head_hunting=&mainSearch=n"

// var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"
// var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)
	for i := 1; i <= totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	oldPageStr := "recruitPage=1"
	newPageStr := fmt.Sprintf("recruitPage=%d", page)
	newURL := strings.Replace(baseURL, oldPageStr, newPageStr, 1)
	fmt.Println(newURL)
	res, err := http.Get(newURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, s *goquery.Selection) { // s는 각 채용공고 카드를 의미함.

	})
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) { // func(i int, s *goquery.Selection) 는 Each 메소드에서 사용해야 하는 형식임. Each() -> ()를 비운 상태에서 에러가 알려줌.
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
