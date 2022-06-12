package main

import (
	"fmt"
	"github.com/notpop/url_getter/config"
	"github.com/notpop/url_getter/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	response, error := http.Get(config.Config.TargetUrl)
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	document, error := goquery.NewDocument(config.Config.TargetUrl)
	if error != nil {
		fmt.Print("connection is failed")
	}
	body, error := document.Find("body").Html()
	if error != nil {
		fmt.Print("document body get failed")
	}
	ioutil.WriteFile(config.Config.GetHtmlPath, []byte(body), os.ModePerm)

	html, _ := ioutil.ReadFile(config.Config.GetHtmlPath)
	stringReader := strings.NewReader(string(html))
	doc, error := goquery.NewDocumentFromReader(stringReader)
	if error != nil {
		fmt.Print("html load failed")
	}

	doc.Find(".article-title-outer > a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		targetUrl := models.NewTargetUrl(url, config.Config.TargetUrl)
		if !models.IsTargetUrl(url) {
			targetUrl.Create()
		}
	})
}
