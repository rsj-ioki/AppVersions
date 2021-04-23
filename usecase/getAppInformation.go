package usecase

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"sort"
)

const androidUrlPrefix = "https://play.google.com/store/apps/details?id="
const iosUrlPrefix = "https://apps.apple.com/de/app/"

type AppsInformation struct {
	AndroidApps []App
	IosApps     []App
}

type App struct {
	Id       string
	Name     string
	Version  string
	Rating   string
	Url      string
	ImageSrc string
	Error    bool
}

func GetAppsInformation(androidAppIds []string, iosAppIds []string) AppsInformation {
	androidAppsChannel := make(chan []App)
	iosAppsChannel := make(chan []App)
	go appInformation(androidAppIds, androidAppsChannel)
	go appInformation(iosAppIds, iosAppsChannel)
	androidApps, iosApps := <-androidAppsChannel, <-iosAppsChannel
	return AppsInformation{
		AndroidApps: androidApps,
		IosApps:     iosApps,
	}
}

func appInformation(appIds []string, appsChannel chan []App) {
	apps := []App{}
	appChannel := make(chan App)
	for _, appId := range appIds {
		go func(appId string) {
			appChannel <- androidAppInfo(appId)
		}(appId)
	}
	for range appIds {
		apps = append(apps, <-appChannel)
	}
	sort.Slice(apps, func(i, j int) bool { return apps[i].Name < apps[j].Name })
	appsChannel <- apps
}

func extractInformation(body []byte, htmlClass string, selector func(int, *goquery.Selection) string) (string, bool) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		return "", false
	}
	selectorResult := ""
	doc.Find(htmlClass).Each(func(i int, s *goquery.Selection) {
		selectorResult = selector(i, s)
	})
	if selectorResult == "" {
		log.Println("selectorResult is empty. Wrong selector?!")
	}
	return selectorResult, true
}
