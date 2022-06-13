package crawl

import (
	"errors"
	"fmt"
	"github.com/notpop/url_getter/config"
	"io"
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

func GetResponseByTargetUrl(targetUrl string) http.Response {
	response, error := http.Get(targetUrl)
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()

	return *response
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
	os.WriteFile(htmlPath, []byte(body), os.ModePerm)

	return true
}

func GetDocumentByHtmlPath(htmlPath string) (*goquery.Document, error) {
	html, _ := os.ReadFile(config.Config.GetHtmlPath)
	stringReader := strings.NewReader(string(html))
	document, error := goquery.NewDocumentFromReader(stringReader)
	if error != nil {
		return nil, errors.New("html load failed")
	}
	return document, nil
}

func SaveImageByTargetUrlDirectoryPathFilePath(targetUrl, directoryPath, filePath string) bool {
	statusCode, status := CheckStatusByTargetUrl(targetUrl)
	if statusCode != 200 {
		log.Fatalf("status code error: %d %s", statusCode, status)
		return false
	}

	err := os.MkdirAll(directoryPath, 0777)
	if err != nil {
		return false
	}

	file, err := os.Create(directoryPath + filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	response := GetResponseByTargetUrl(targetUrl)
	io.Copy(file, response.Body)
	return true
}
