package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/artemkakun/skillshot-tgbot/structs"
	"log"
	"net/http"
	"strings"
)

func GetVacanciesLinksList() (vacancies []string) {
	siteData, err := http.Get("https://www.skillshot.pl/")
	if err != nil {
		log.Fatal(err)
	}

	if siteData.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", siteData.StatusCode, siteData.Status)
	}

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

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	} else {
		body_data, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		one_vacancy_data.Link = full_vacancy_link
		one_vacancy_data.Title = body_data.Find("h1").Text()
		one_vacancy_data.Employer = body_data.Find("b a").Text()
		one_vacancy_data.WorkPlace = getVacancyLocation(body_data.Find("p").First().Text())
		one_vacancy_data.WorkType = body_data.Find(".badge-job-category").Text()

		if body_data.Find(".badge-job-stala").Text() != "" {
			one_vacancy_data.EmploymentType = body_data.Find(".badge-job-stala").Text()
		} else if body_data.Find(".badge-job-zlecenie").Text() != "" {
			one_vacancy_data.EmploymentType = body_data.Find(".badge-job-zlecenie").Text()
		} else {
			one_vacancy_data.EmploymentType = body_data.Find(".badge-job-inne").Text()
		}

		//TODO - Need to create algorithm for filtering this data
		//one_vacancy_data.Text = getVacancyLocation(body_data.Find("p").Text())
	}
	return one_vacancy_data
}

func getVacancyLocation(unknown_location string) (location string) {
	split_string := strings.Fields(unknown_location)

	for i, one_word := range split_string {
		if one_word == "w" {
			location_buffer := split_string[i+1:]

			for _, one_location_word := range location_buffer {
				location += one_location_word + " "
			}
		}
	}

	return location
}
