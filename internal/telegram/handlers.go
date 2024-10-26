package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
)

func (t *Telegram) MainHandler(c echo.Context) error {
	log.Printf("MainHandler")
	var update tgbotapi.Update

	err := c.Bind(&update)
	if err != nil {
		log.Print("Cannot bind update", err)
		return c.JSON(204, nil)
	}
	if update.Message != nil {
		log.Printf("update: %+v", update)
	}
	// 204 err if nil

	if update.Message.Text != "" {
		log.Printf("textcontets: %s from %s", update.Message.Text, update.Message.From)
		err := t.defaultAnswer(update)
		if err != nil {
			return c.JSON(204, nil)
		}
	}
	return c.JSON(200, nil)
}

func (t *Telegram) defaultAnswer(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.From.ID, "rrrra!")

	msg.ReplyToMessageID = update.Message.MessageID
	t.Bot.Send(msg)
	return nil
}
