package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/go/classroom-8/students/workshop-6/pkg/kafka"
)

type Message struct {
	ID    int
	Value string
}

const (
	topicName = "example-topic"
)

var brokers = []string{
	"127.0.0.1:9091",
	"127.0.0.1:9092",
	"127.0.0.1:9093",
}

func main() {
	producer, err := kafka.NewAsyncProducer(brokers,
		kafka.WithRequiredAcks(sarama.NoResponse),
		kafka.WithProducerPartitioner(sarama.NewHashPartitioner),
		kafka.WithMaxOpenRequests(5),
		kafka.WithMaxRetries(5),
		kafka.WithRetryBackoff(10*time.Millisecond),
		kafka.WithProducerFlushMessages(3),
		kafka.WithProducerFlushFrequency(5*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer producer.Close() // Не забываем освобождать ресурсы :)

	m := Message{
		ID:    1,
		Value: "example",
	}
	send(producer, m)

	m.ID = 2
	send(producer, m)
	fmt.Println("messages send async...")

	time.Sleep(5 * time.Second)
	fmt.Println("5 sec left")

	// ... сообщения отправятся

	time.Sleep(1 * time.Second)
}

func send(producer sarama.AsyncProducer, m Message) {
	bytes, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	msg, err := kafka.BuildMessage(topicName, fmt.Sprint(m.ID), bytes, "x-header-example", "example-header-value")
	if err != nil {
		log.Fatal(err)
	}

	producer.Input() <- msg
}
