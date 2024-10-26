package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"github.com/randomtoy/go-smallgptbot/internal/telegram"
)

func main() {
	// Happiness_bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	whInfo, _ := bot.GetWebhookInfo()
	log.Printf("whInfo: %#v\n", whInfo)
	t := telegram.New(bot)
	log.Printf("telegram: %+v\n", t)

	e := echo.New()
	e.POST("/", t.MainHandler)
	e.Start(":" + os.Getenv("PORT"))
}
