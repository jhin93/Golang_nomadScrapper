package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
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

// var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"
// var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

// Scrape Site by a term
func Scrape(term string) {
	var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=" + term + "&recruitPage=1&recruitSort=relation&recruitPageCount=40&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10&show_applied=&quick_apply=&except_read=&ai_head_hunting=&mainSearch=n"
	// 1. 총 페이지 수를 가져온다(getPages 함수).
	// 2. 각 페이지별로 goroutine을 생성(getPage, getPage, getPage...)
	// 3. 각 getPage는 각 일자리정보 별로 goroutine을 생성.
	// 4. 정리 : 기존 방식(1페이지에서 순서대로 50개 정보 취합 -> 2페이지 순서대로 ...)에서 goroutine 방식으로 병렬 처리(모든 페이지에서 동시에 일자리 정보를 각각 취합)
	var jobs []extractedJob
	c := make(chan []extractedJob) // 하나의 extractedJob이 아닌 다수의 extractedJob를 담은 slice 여야 함.

	totalPages := getPages(baseURL)
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		// getPage는 1페이지에 담긴 여러 카드의 정보들을 담은 배열을 반환한다.
		go getPage(i, baseURL, c) // getPage에 데이터를 받아올 채널(c)를 인자로 보낸다.
	}

	for i := 0; i < totalPages; i++ { // 위 반복문과 같은 맥락으로 반복문 작성
		extractedJobs := <-c // 메세지를 기다리다가 채널에 메세지가 전송되면 extractedJob에 저장
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs) // 모든 url의 결과물을 jobs(func main의 jobs)에 담고 반복문이 끝난 다음 맨 마지막에 결과를 도출한다.
	fmt.Println("Done, extracted", len(jobs))
}

func getPage(page int, url string, mainC chan<- []extractedJob) { // 인자 mainC는 데이터를 받기만 할 것이기에 chan<-라고 입력
	var jobs []extractedJob      // 각각의 정보 카드(extractedJob) 으로 이루어진 slice 타입의 변수 jobs
	c := make(chan extractedJob) // 이 채널(변수 c를 의미)에 전송할 값은 extractedJob 타입이다.
	oldPageStr := "recruitPage=1"
	newPageStr := fmt.Sprintf("recruitPage=%d", page)
	newURL := strings.Replace(url, oldPageStr, newPageStr, 1)
	fmt.Println(newURL)
	res, err := http.Get(newURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection) { // s는 각 채용공고 카드를 의미함.
		go extractJob(card, c) // extractJob 함수에 채널(위의 변수 c)을 인자로 입력.
	})

	for i := 0; i < searchCards.Length(); i++ { // 페이지 장 수만큼 반복문 돔
		job := <-c               // 채널에 받은 데이터를 변수 job에 저장
		jobs = append(jobs, job) // slice 변수 jobs에 데이터 job 저장.
	}

	mainC <- jobs // jobs를 리턴하는 대신 메인함수의 채널 mainC에 데이터 저장
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
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

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv") // 파일 이름 jobs.csv
	checkErr(err)

	w := csv.NewWriter(file) // 파일을 새로 만드는 것
	defer w.Flush()          // Flush()는 함수가 끝나는 시점에 파일에 데이터를 입력하는 함수

	headers := []string{"ID", "Title", "Location", "Career", "Sector"} // headers는 함수의 윗부분

	wErr := w.Write(headers) // Write 함수로 headers를 w 에 작성.
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.location, job.career, job.sector}
		jwErr := w.Write(jobSlice)
		checkErr((jwErr)) // 여기까지 끝나고 나면 defer w.Flush가 발동해서 데이터가 파일에 입력될 것.
	}
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

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	title := cleanString(card.Find(".area_job>h2").Text())
	location := cleanString(card.Find(".area_job>.job_condition span:first-child").Text())
	career := cleanString(card.Find(".area_job>.job_condition span:second-child").Text())
	sector := cleanString(card.Find(".area_job>.job_sector").Text())
	c <- extractedJob{ // 채널에 값 받기
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
