package main

import (
	"github.com/notpop/url_getter/config"
	"github.com/notpop/url_getter/crawl"
	"github.com/notpop/url_getter/models"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	statusCode, status := crawl.CheckStatusByTargetUrl(config.Config.TargetUrl)
	if statusCode != 200 {
		log.Fatalf("status code error: %d %s", statusCode, status)
	}

	if !crawl.SaveHtmlByTargetUrl(config.Config.TargetUrl, config.Config.GetHtmlPath) {
		log.Fatal("could not save html")
	}

	document, error := crawl.GetDocumentByHtmlPath(config.Config.GetHtmlPath)
	if error != nil {
		log.Println(error)
	}

	document.Find(config.Config.OriginSourceSelector).Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		targetUrl := models.NewTargetUrl(url, config.Config.TargetUrl)
		if !models.IsTargetUrl(url) {
			targetUrl.Create()
		}
	})

	dfTargetUrl, error := models.GetAllTargetUrl(config.Config.SearchLimit)
	if error != nil {
		log.Println(error)
	}

	for i, url := range dfTargetUrl.Urls() {
		statusCode, status := crawl.CheckStatusByTargetUrl(url)
		if statusCode != 200 {
			log.Fatalf("status code error: %d %s", statusCode, status)
		}

		temporaryHtmlPath := "./htmls/" + strconv.Itoa(i) + ".html"
		if !crawl.SaveHtmlByTargetUrl(url, temporaryHtmlPath) {
			log.Fatal("could not save html")
		}

		document, error := crawl.GetDocumentByHtmlPath(temporaryHtmlPath)
		if error != nil {
			log.Println(error)
		}

		document.Find(config.Config.SubSelector).Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Attr("src")
			// 新規テーブルに保存
			targetUrl := models.NewTargetUrl(url, config.Config.TargetUrl)
			if !models.IsTargetUrl(url) {
				targetUrl.Create()
			}
		})
	}
}
