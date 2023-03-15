package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"telegram_bot_golang/pkg/models"
	"time"
)

func (b *Bot) handleCommand(message *tgbotapi.Message, updates tgbotapi.UpdatesChannel) error {
	switch message.Command() {
	case viper.GetString("commands.start"):
		return b.handleStartCommand(message)
	case viper.GetString("commands.getFirstQueryInfo"):
		return b.handleGetFirstQueryTimeCommand(message)
	case viper.GetString("commands.getAllQueriesCount"):
		return b.handleGetUserAllQueriesCountCommand(message)
	case viper.GetString("commands.addNewEvent"):
		return b.addNewEvent(message, updates)
	default:
		return ErrorUnknownCommand
	}
}

func (b *Bot) addNewEvent(message *tgbotapi.Message, updates tgbotapi.UpdatesChannel) error {
	logrus.Printf("[%s] %s", message.From.UserName, "Provet")
	msg := tgbotapi.NewMessage(message.Chat.ID, viper.GetString("messages.responses.replyAddNewDate"))
	b.bot.Send(msg)

	var date string
	first := true

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommand(update.Message, updates)
			return nil
		}

		val := strings.Split(update.Message.Text, ".")
		logrus.Printf("%s", update.Message.Text)
		if len(val) == 2 {
			if ch := ValidateDate(val[0], val[1]); ch != true {
				return ErrorWrongDateFormat
				break
			}
			date = update.Message.Text
			first = false
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, viper.GetString("messages.responses.replyAddNewDescription"))
			b.bot.Send(msg)
			continue
		} else if first && len(val) != 2 {
			return ErrorWrongNewDateFormat
			break
		}
		if !first {
			err := b.handleNewEventDate(date, update.Message)
			if err != nil {
				return ErrorUnknown
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, viper.GetString("messages.responses.replySuccessSaveNewDate"))
				b.bot.Send(msg)
			}
			break
		}
	}

	return nil
}

func (b *Bot) handleNewEventDate(date string, message *tgbotapi.Message) error {
	return b.service.CreateDate(&models.Dates{ChatID: message.Chat.ID, Description: message.Text, Date: date})
}

func ValidateDate(month, day string) bool {
	output := false

	monthInt, err := strconv.Atoi(month)
	if err != nil {
		return false
	}
	if monthInt > 0 && monthInt <= 12 {
		output = true
	}

	dayInt, err := strconv.Atoi(day)
	if err != nil {
		return false
	}
	if dayInt > 0 && dayInt <= 31 {
		output = true
	}

	if monthInt == 02 && dayInt >= 30 {
		return false
	}

	return output
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	logrus.Printf("[%s] %s", message.From.UserName, message.Text)

	val := strings.Split(message.Text, ".")

	if len(val) == 3 {
		if ch := ValidateDate(val[1], val[2]); ch != true {
			return ErrorUnknownCommand
		} else {
			response, err := getHolidays(val[0], val[1], val[2])
			if err != nil {
				logrus.Printf("Invalid message %s", err.Error())
				return ErrorWrongDateFormat
			}

			datearray, _ := b.service.GetDateByDate(val[1] + "." + val[2])
			datestring := strings.Join(datearray, ", ")

			if len(response.Holidays) != 0 {
				b.service.CreateUser(models.NewUsers(message.Chat.ID, message.From.UserName, time.Now().String()+" Message: "+message.Text, 1))

				msg := tgbotapi.NewMessage(message.Chat.ID, response.Holidays[0].Name+" \n"+datestring)
				b.bot.Send(msg)
				return nil
			} else if len(datearray) != 0 {
				msg := tgbotapi.NewMessage(message.Chat.ID, datestring)
				b.bot.Send(msg)
				return nil
			} else {
				msg := tgbotapi.NewMessage(message.Chat.ID, viper.GetString("messages.responses.simpleDay")+" \n"+datestring)
				b.bot.Send(msg)
				return nil
			}
		}
	}

	return ErrorWrongDateFormat

}

func (b *Bot) handleGetFirstQueryTimeCommand(message *tgbotapi.Message) error {
	logrus.Printf("[%s] %s", message.From.UserName, message.Text)

	userID, err := b.service.GetUserID(message.Chat.ID)
	if err != nil {
		return NullResponse
	}

	firstQueryTime, err := b.service.GetUserFirstQueryTime(int64(userID))
	if err != nil {
		return NullResponse
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, firstQueryTime)
	_, err = b.bot.Send(msg)
	return nil
}

func (b *Bot) handleGetUserAllQueriesCountCommand(message *tgbotapi.Message) error {
	logrus.Printf("[%s] %s", message.From.UserName, message.Text)

	allQueriesCount, err := b.service.GetUserAllQueriesCount(message.Chat.ID)
	if err != nil {
		return NullResponse
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, strconv.FormatInt(int64(allQueriesCount), 10))
	_, err = b.bot.Send(msg)
	return nil
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	logrus.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, viper.GetString("messages.responses.start"))
	_, err := b.bot.Send(msg)
	return err
}
