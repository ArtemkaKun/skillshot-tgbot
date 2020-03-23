package bot

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

var bot *tgbotapi.BotAPI

func init() {
	bot = botStart()
}

func botStart() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI("BOT TOKEN HERE")
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
			msg.Text = "value"
		}

		_, err := my_bot.Send(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
