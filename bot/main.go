package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/rAndrade360/go-link-shortener/bot/adapter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Err to get Telegram Token ", err.Error())
	}

	api := adapter.NewAdapterApi("localhost:8081")
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal("Err to get Telegram Token ", err.Error())
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 20

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/start" {
			txt := "Welcome! This is your url shortener bot. To short an url, send te following message:\n\n short <your url>"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, txt)

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}

		if strings.Contains(update.Message.Text, "short") {
			url := strings.Split(update.Message.Text, " ")[1]

			url_short, err := api.ShortUrl(url)
			if err != nil {
				log.Fatal("Error to Call Short", err.Error())
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, url_short)

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}

	}

	url_short, err := api.ShortUrl("http://github.com/rAndrade360/GRPC")
	if err != nil {
		log.Fatal("Error to Call gRPC", err.Error())
	}

	log.Println("Returned: ", url_short)
}
