package parser

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

var siteData *http.Response

func init() {
	var err error

	siteData, err = http.Get("https://www.skillshot.pl/")
	if err != nil {
		log.Fatal(err)
	}

	if siteData.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", siteData.StatusCode, siteData.Status)
	}
}

func GetVacanciesLinksList() (vacancies []string) {
	defer siteData.Body.Close()

	body_data, err := goquery.NewDocumentFromReader(siteData.Body)
	if err != nil {
		log.Fatal(err)
	}

	body_data.Find("td").Each(func(_ int, vacancies_table *goquery.Selection) {
		vacancies_links := vacancies_table.Find("a")
		for _, one_link := range vacancies_links.Nodes {
			vacancies = append(vacancies, one_link.Attr[0].Val)
		}
	})

	return vacancies
}
