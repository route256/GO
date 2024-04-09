package worker

import (
	"context"
	"testing"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/worker/mocks"
)

func TestMessageListenerWorkerRun(t *testing.T) {
	t.Run("after interrupting context processing new messages will stop", func(t *testing.T) {
		t.Parallel()

		messageFetcher := mocks.NewMockMessageFetcher(gomock.NewController(t))
		messageProcessor := mocks.NewMockMessageProcessor(gomock.NewController(t))

		chWithUpdates := make(chan tgbotapi.Update, 100)
		messageProcessor.EXPECT().IncomingMessage(gomock.Any(), gomock.Any()).Return(nil).Times(7)

		messageFetcher.EXPECT().Stop().Times(1)
		messageFetcher.EXPECT().Start().Return(chWithUpdates).Times(1)

		worker := NewMessageListenerWorker(messageFetcher, messageProcessor)

		ctx, cancel := context.WithCancel(context.TODO())

		// генерация сообщений
		go func(ch chan<- tgbotapi.Update) {
			for i := 0; i < 10; i++ {
				// эмулируем задержку публикации сообщений
				time.Sleep(50 * time.Millisecond)
				if i == 7 {
					cancel()
				}

				ch <- tgbotapi.Update{Message: &tgbotapi.Message{From: &tgbotapi.User{ID: 1230}, Text: "/command"}}
			}

			close(chWithUpdates)
		}(chWithUpdates)

		worker.Run(ctx)

		assert.Error(t, ctx.Err())
	})
}
