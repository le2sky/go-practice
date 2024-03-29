package scrapper

import (
	"fmt"
	"log"
  "strings"
	"net/http"
	"strconv"
  "encoding/csv"
  "os"
	"github.com/PuerkitoBio/goquery"
)


type extractedJob struct {
  id string
  location string
  title string
  summary string
}

// Scraper
func Scrape (term string){
 var baseURL string = "https://kr.indeed.com/jobs?q="+ term +"&limit=50"
 var jobs []extractedJob
 c:= make(chan []extractedJob)
 totalPages := getPages(baseURL)


 for i:=0; i< totalPages; i++{
    go getPage(i, baseURL , c)
 }

  for i:=0; i<totalPages; i++{
    extractedJob := <- c
    jobs = append(jobs, extractedJob...)
  }

 writeJobs(jobs)
 fmt.Println("Done, extracted:", len(jobs))
}


func getPage(page int, url string,mainC chan<- []extractedJob){
  var jobs []extractedJob
  c := make(chan extractedJob)
  pageURL := url + "&start=" + strconv.Itoa(page*50)
  fmt.Println("Requesting ", pageURL)
  res, err := http.Get(pageURL)
  checkErr(err)
  checkCode(res)

  defer res.Body.Close()

  doc, err := goquery.NewDocumentFromReader(res.Body)
  checkErr(err)

  searchCards := doc.Find(".tapItem")
  searchCards.Each(func(i int, card *goquery.Selection) {
    go extractJob(card, c)
  })
  for i:=0; i<searchCards.Length() ;i++{
    job := <- c
    jobs = append(jobs, job)
  }
  mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
    id, _ := card.Attr("data-jk")
    title := CleanString(card.Find("h2>span").Text())
    location := CleanString(card.Find(".companyLocation").Text())
    summary := CleanString(card.Find(".job-snippet").Text())

    c <- extractedJob{
      id: id, 
      title: title, 
      location: location, 
      summary: summary,
    }
}


func getPages(url string) int {
  pages := 0

  res, err := http.Get(url)
  checkErr(err)
  checkCode(res)
  
  defer res.Body.Close()

  doc, err := goquery.NewDocumentFromReader(res.Body)
  checkErr(err)

  doc.Find(".pagination").Each(func(i int, s *goquery.Selection){
    pages = s.Find("a").Length()
  })

  return pages
}

func checkErr(err error){
  if err != nil {
    log.Fatalln(err)
  } 
}

func checkCode(res *http.Response){
  if res.StatusCode != 200 {
    log.Fatalln("Request failed with Status: ", res.StatusCode)
  }
}

func CleanString(str string) string {
  return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func writeJobs(jobs []extractedJob){
  file, err := os.Create("jobs.csv")
  checkErr(err)
  w := csv.NewWriter(file)
  utf8bom := []byte{0xEF, 0xBB, 0xBF}
  file.Write(utf8bom)
  defer w.Flush()

  headers := []string{"Id", "Title", "Location", "Summary"}

  wErr := w.Write(headers)
  checkErr(wErr)
  
  for _, job := range jobs {
    jobSlice := []string{
      job.id,
      job.title,
      job.location,
      job.summary,
    }
    jwErr := w.Write(jobSlice)
    checkErr(jwErr)
  }

}