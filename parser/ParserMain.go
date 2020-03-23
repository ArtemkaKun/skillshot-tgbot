package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/artemkakun/skillshot-tgbot/structs"
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

func GetVacancyData(vacancy_link string) (one_vacancy_data structs.VacancyData) {
	full_vacancy_link := "https://www.skillshot.pl" + vacancy_link

	res, err := http.Get(full_vacancy_link)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if siteData.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	} else {
		body_data, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		one_vacancy_data.Title = body_data.Find("h1").Text()
		one_vacancy_data.Employer = body_data.Find("b a").Text()
		/*body_data.Find("td").Each(func(_ int, vacancies_table *goquery.Selection) {
			vacancies_links := vacancies_table.Find("a")
			for _, one_link := range vacancies_links.Nodes {
				vacancies = append(vacancies, one_link.Attr[0].Val)
			}
		})*/
	}

	return one_vacancy_data
}
