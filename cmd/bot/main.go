package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"telegram_bot_golang/pkg/repositories"
	"telegram_bot_golang/pkg/services"
	"telegram_bot_golang/pkg/telegram"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := repositories.NewPostgresDB(repositories.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Printf("Failed to initializ db: %s", err.Error())
	}

	repos := repositories.NewRepository(db)
	services := services.NewService(repos)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_KEY"))
	if err != nil {
		logrus.Fatalf("failed NewBotAPI: %s", err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot, services)
	if err := telegramBot.Start(); err != nil {
		logrus.Fatal(err)
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
