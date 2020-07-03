package usecase

import (
	"net/http"
	"log"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func androidVersionForAppId(androidAppId string) string {
	url := androidUrlPrefix + androidAppId
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	version := ""
	doc.Find(".hAyfc .htlgb").Each(func(i int, s *goquery.Selection) {
		if i == 6 {
			version = s.Text()
		}
	})
	return version
}

func iosVersionForAppId(iosAppId string) string {
	url := iosUrlPrefix + iosAppId
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	version := ""
	doc.Find(".whats-new__latest__version").Each(func(i int, s *goquery.Selection) {
		version = strings.Replace(s.Text(), "Version ", "" , -1)
	})
	return version
}
