package telegram

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"github.com/randomtoy/go-smallgptbot/internal/openai"
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
	oa := openai.New(os.Getenv("OPENAI_APITOKEN"))
	oa.Model = "gpt-3.5-turbo"
	oa.System = "professional teacher who knows TOEFL preparation perfectly"
	oa.User = update.Message.Text

	resp, err := oa.Send()
	log.Printf("response: %+v", resp)
	if err != nil {
		return err
	}

	for i := range resp.Choises {
		msg := tgbotapi.NewMessage(update.Message.From.ID, resp.Choises[i].Message.Content)
		msg.ReplyToMessageID = update.Message.MessageID
		t.Bot.Send(msg)
	}

	return nil
}
