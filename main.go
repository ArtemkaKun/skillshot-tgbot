package main

import (
	"github.com/artemkakun/skillshot-tgbot/parser"
)

func main() {
	vacancies_list := parser.GetVacanciesLinksList()

	//var wg sync.WaitGroup
	//wg.Add(len(vacancies_list))

	for _, one_vacancy_link := range vacancies_list {
		parser.GetVacancyData(one_vacancy_link)
	}

	//wg.Wait()
}
