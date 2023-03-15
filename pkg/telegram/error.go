package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
)

var (
	ErrorUnknownCommand     = errors.New("Неизвестная команда!")
	ErrorWrongDateFormat    = errors.New("Неправильный формат даты!")
	ErrorWrongNewDateFormat = errors.New("Неправильный формат новой даты!")
	ErrorUnknown            = errors.New("Чтоөто пошло не так!")
	NullResponse            = errors.New("База пустая")
)

func (b *Bot) handleError(chatID int64, err error) {
	var messageText string

	switch err {
	case ErrorUnknownCommand:
		messageText = viper.GetString("messages.errors.unknownCommand")
	case ErrorWrongDateFormat:
		messageText = viper.GetString("messages.errors.wrongDateFormat")
	case ErrorWrongNewDateFormat:
		messageText = viper.GetString("messages.errors.wrongNewDateFormat")
	case ErrorUnknown:
		messageText = viper.GetString("messages.errors.unknownError")
	case NullResponse:
		messageText = viper.GetString("messages.errors.nullResponse")
	default:
		messageText = viper.GetString("messages.errors.default")
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	b.bot.Send(msg)
}
