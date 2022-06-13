package crawl

import (
	"errors"
	"fmt"
	"github.com/notpop/url_getter/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CheckStatusByTargetUrl(targetUrl string) (statusCode int, status string) {
	response, error := http.Get(targetUrl)
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()

	return response.StatusCode, response.Status
}

func SaveHtmlByTargetUrl(targetUrl, htmlPath string) bool {
	document, error := goquery.NewDocument(config.Config.TargetUrl)
	if error != nil {
		fmt.Println("connection is failed")
		return false
	}
	body, error := document.Find("body").Html()
	if error != nil {
		fmt.Println("document body get failed")
		return false
	}
	ioutil.WriteFile(htmlPath, []byte(body), os.ModePerm)

	return true
}

func GetDocumentByHtmlPath(htmlPath string) (*goquery.Document, error) {
	html, _ := ioutil.ReadFile(config.Config.GetHtmlPath)
	stringReader := strings.NewReader(string(html))
	document, error := goquery.NewDocumentFromReader(stringReader)
	if error != nil {
		return nil, errors.New("html load failed")
	}
	return document, nil
}
