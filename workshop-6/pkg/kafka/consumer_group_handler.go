package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

var _ sarama.ConsumerGroupHandler = (*ConsumerGroupHandler)(nil)

type ConsumerGroupHandler struct {
	ready chan bool
}

func NewConsumerGroupHandler(
// map[topic]TopichHndler
) *ConsumerGroupHandler {
	return &ConsumerGroupHandler{
		ready: make(chan bool),
	}
}

func (h *ConsumerGroupHandler) Ready() <-chan bool {
	return h.ready
}

// Setup Начинаем новую сессию, до ConsumeClaim
func (h *ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	close(h.ready)

	return nil
}

// Cleanup завершает сессию, после того, как все ConsumeClaim завершатся
func (h *ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim читаем до тех пор пока сессия не завершилась
func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			log.Printf("Message claimed: value = %v, timestamp = %v, topic = %s",
				string(message.Value),
				message.Timestamp,
				message.Topic,
			)

			// коммит сообщения "руками"
			//
			session.MarkMessage(message, "")
			//
		case <-session.Context().Done():
			return nil
		}
	}
}
