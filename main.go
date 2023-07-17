package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	career   string
	sector   string
}

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

	searchCards.Each(func(i int, card *goquery.Selection) { // s는 각 채용공고 카드를 의미함.
		extractJob(card)
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

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) { // func(i int, s *goquery.Selection) 는 Each 메소드에서 사용해야 하는 형식임. Each()까지만 쓰고 ()를 비운 상태에서 에러가 알려줌.
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

func extractJob(card *goquery.Selection) {
	id, _ := card.Attr("value")
	title := cleanString(card.Find(".area_job>h2").Text())
	location := cleanString(card.Find(".area_job>.job_condition span:first-child").Text())
	career := cleanString(card.Find(".area_job>.job_condition span:second-child").Text())
	sector := cleanString(card.Find(".area_job>.job_sector").Text())
	fmt.Println(id, title, location, career, sector)
}

func cleanString(str string) string {
	// 양쪽 끝 공백은 TrimSpace가 지우고, Fields는 내부 공백을 지운 후 배열에 단어를 담는다. 이후 배열에 담긴 단어들을 join으로 " "로 합친다.
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
