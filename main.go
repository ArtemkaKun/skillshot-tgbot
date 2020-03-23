package main

import (
	"fmt"
	"github.com/artemkakun/skillshot-tgbot/parser"
	"github.com/artemkakun/skillshot-tgbot/structs"
	"time"
)

func main() {
	go parserLoop()
}

func parserLoop() {
	var actual_vacancies_list []string
	var previous_vacancies_list []string

	for true {
		actual_vacancies_list = parser.GetVacanciesLinksList()

		new_vacancies := subtractSlices(actual_vacancies_list, previous_vacancies_list)

		if len(new_vacancies) > 0 {
			fmt.Println("New vaca")
			for _, one_vacancy_link := range new_vacancies {
				vacancy_data := parser.GetVacancyData(one_vacancy_link)
				//send bot req
			}
		}

		time.Sleep(1 * time.Minute)
	}
}

func subtractSlices(A, B []string) (diff_slice []string) {
	//A - B
	for _, one_link_A := range A {
		contain := false

		for _, one_link_B := range B {
			if one_link_A == one_link_B {
				contain = true
				break
			}
		}

		if contain {
			continue
		}

		diff_slice = append(diff_slice, one_link_A)
	}

	return diff_slice
}
