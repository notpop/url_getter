package main

import (
  "fmt"
	"io/ioutil"
  "log"
  "net/http"
	"os"
	"strings"

  "github.com/PuerkitoBio/goquery"
)

const TARGET_URL = "https://www.youtube.com/"
// const TARGET_URL = "https://www.youtube.com/watch?v=dQAGDdkW4ag"

const TEMPORARY_HTML_FILE_PATH = "./htmls/target.html"

func main() {
  response, error := http.Get(TARGET_URL)
  if error != nil {
    log.Fatal(error)
  }
  defer response.Body.Close()
  if response.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
  }

	document, error := goquery.NewDocument(TARGET_URL)
	if error != nil {
			fmt.Print("connection is failed")
	}
	body, error := document.Find("body").Html()
	if error != nil {
			fmt.Print("document body get failed")
	}
	ioutil.WriteFile(TEMPORARY_HTML_FILE_PATH, []byte(body), os.ModePerm)

	html, _ := ioutil.ReadFile(TEMPORARY_HTML_FILE_PATH)
	stringReader := strings.NewReader(string(html))
	doc, error := goquery.NewDocumentFromReader(stringReader)
	if error != nil {
			fmt.Print("html load failed")
	}
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			fmt.Println(url)
	})
}
