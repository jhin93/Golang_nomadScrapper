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
	var jobs []extractedJob
	totalPages := getPages()
	fmt.Println(totalPages)
	for i := 1; i <= totalPages; i++ {
		extractedJobs := getPage(i)           // getPage는 1페이지에 담긴 여러 카드의 정보들을 담은 배열을 반환한다.
		jobs = append(jobs, extractedJobs...) // append()를 통해 배열들 각각의 contents 들을 또다시 합친다. extractedJobs의 contents 다수를 표현하는 것이 바로 '...' 이다. 즉, 배열들 각각의 contents들을 뽑아내어 합치는 것.
	}
	fmt.Println(jobs) // 모든 url의 결과물을 jobs(func main의 jobs)에 담고 반복문이 끝난 다음 맨 마지막에 결과를 도출한다.
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob // 각각의 정보 카드(extractedJob) 으로 이루어진 slice 타입의 변수 jobs
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
		job := extractJob(card)
		jobs = append(jobs, job)
	})
	return jobs // 다수의 job 을 담은 배열 'jobs' 를 반환.
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

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("value")
	title := cleanString(card.Find(".area_job>h2").Text())
	location := cleanString(card.Find(".area_job>.job_condition span:first-child").Text())
	career := cleanString(card.Find(".area_job>.job_condition span:second-child").Text())
	sector := cleanString(card.Find(".area_job>.job_sector").Text())
	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		career:   career,
		sector:   sector,
	}
}

func cleanString(str string) string {
	// 양쪽 끝 공백은 TrimSpace가 지우고, Fields는 내부 공백을 지운 후 배열에 단어를 담는다. 이후 배열에 담긴 단어들을 join으로 " "로 합친다.
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
