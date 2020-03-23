package bot

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/artemkakun/skillshot-tgbot/db"
	"github.com/artemkakun/skillshot-tgbot/structs"
	"log"
)

var bot *tgbotapi.BotAPI

func init() {
	bot = botStart()
}

func botStart() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI("1126466003:AAELrk2wtxqzuneZkbyxsZna62XsJqHEt-4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Autorised on account %s", bot.Self.UserName)

	return bot
}

func BotUpdateLoop(my_bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := my_bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {

		if update.Message == nil {
			continue
		}

		chat_id := update.Message.Chat.ID
		msg := tgbotapi.NewMessage(chat_id, "")

		switch update.Message.Command() {
		case "start":
			msg.Text = "Now you will be alert about new vacancies from Skillshot.pl"
			db.AddNewUser(chat_id)
		}

		_, err := my_bot.Send(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func SendNewVacancy(vacancy_data structs.VacancyData) {
	users_ids := db.GetUsers()

	for _, one_id := range users_ids {
		msg := tgbotapi.NewMessage(one_id, fmt.Sprintf("%v\nPracowadwca: #%v\nMiejscowość: #%v\nRodzaj pracy: #%v\nTyp umowy: #%v\nLink: %v\n", vacancy_data.Title, vacancy_data.Employer, vacancy_data.WorkPlace, vacancy_data.WorkType, vacancy_data.EmploymentType, vacancy_data.Link))
		_, err := bot.Send(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func GetBot() (my_bot *tgbotapi.BotAPI) {
	my_bot = bot

	return my_bot
}
