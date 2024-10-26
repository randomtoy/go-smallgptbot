package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	Bot *tgbotapi.BotAPI
}

func New(bot *tgbotapi.BotAPI) *Telegram {
	return &Telegram{
		Bot: bot,
	}
}
