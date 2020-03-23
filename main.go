package main

import (
	"fmt"
	"github.com/artemkakun/skillshot-tgbot/parser"
)

func main() {
	vacancies_list := parser.GetVacanciesLinksList()

	for _, one_vacancy_text := range vacancies_list {
		fmt.Println(one_vacancy_text)
	}
}
