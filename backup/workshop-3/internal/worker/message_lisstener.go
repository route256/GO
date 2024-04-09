package worker

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/model/messages"
)

type MessageFetcher interface {
	Start() tgbotapi.UpdatesChannel
	Request(callback tgbotapi.CallbackConfig) error
	Stop()
}

type MessageProcessor interface {
	IncomingMessage(ctx context.Context, msg messages.Message) error
}

type MessageListenerWorker struct {
	messageFetcher MessageFetcher
	processor      MessageProcessor
}

func NewMessageListenerWorker(messageFetcher MessageFetcher, processor MessageProcessor) *MessageListenerWorker {
	return &MessageListenerWorker{
		messageFetcher: messageFetcher,
		processor:      processor,
	}
}

func (worker *MessageListenerWorker) Run(ctx context.Context) {
	for update := range worker.messageFetcher.Start() {
		select {
		// По заданию нужно отменять операцию на основании контекста
		case <-ctx.Done():
			worker.messageFetcher.Stop()
			log.Println("stopped receiving new message from tg bot")
			return
		default:
			if err := worker.processing(ctx, update); err != nil {
				log.Println(err)
			}
		}
	}
}

func (worker *MessageListenerWorker) processing(ctx context.Context, update tgbotapi.Update) error {
	if update.Message != nil { // If we got a message
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		return worker.processor.IncomingMessage(ctx, messages.Message{
			Text:   update.Message.Text,
			UserID: update.Message.From.ID,
		})
	} else if update.CallbackQuery != nil {
		// Respond to the callback query, telling Telegram to show the user
		// a message with the data received.
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
		if err := worker.messageFetcher.Request(callback); err != nil {
			log.Println("error processing message:", err)
		}

		err := worker.processor.IncomingMessage(ctx, messages.Message{
			Text:   update.CallbackQuery.Data,
			UserID: update.CallbackQuery.From.ID,
		})
		if err != nil {
			log.Println("error processing message:", err)
		}

		// And finally, send a message containing the data received.
		//msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
		//if _, err := c.client.Send(msg); err != nil {
		//	panic(err)
		//}
	}

	return nil
}
