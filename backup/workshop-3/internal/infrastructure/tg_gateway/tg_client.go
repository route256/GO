package tg_gateway

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type TokenGetter interface {
	Token() string
}

type Client struct {
	client *tgbotapi.BotAPI
}

func New(config TokenGetter) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(config.Token())
	if err != nil {
		return nil, errors.Wrap(err, "NewBotAPI")
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) SendMessage(text string, userID int64, keyboardRows ...map[string]string) error {
	msg := tgbotapi.NewMessage(userID, text)

	if len(keyboardRows) != 0 {
		var rows [][]tgbotapi.InlineKeyboardButton

		for _, buttons := range keyboardRows {
			var row []tgbotapi.InlineKeyboardButton
			for text, data := range buttons {
				row = append(row, tgbotapi.NewInlineKeyboardButtonData(text, data))
			}

			rows = append(rows, tgbotapi.NewInlineKeyboardRow(row...))
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	}

	_, err := c.client.Send(msg)
	if err != nil {
		return errors.Wrap(err, "client.Send")
	}
	return nil
}

func (c *Client) Start() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return c.client.GetUpdatesChan(u)
}

func (c *Client) Stop() {
	c.client.StopReceivingUpdates()
}

func (c *Client) Request(callback tgbotapi.CallbackConfig) error {
	_, err := c.client.Request(callback)
	return err
}
