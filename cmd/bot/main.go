package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"telegram_bot_golang/pkg/telegram"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6048429033:AAHnsVwcrAmjoIiFLuxtT64z8_vc_5lhkO4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
