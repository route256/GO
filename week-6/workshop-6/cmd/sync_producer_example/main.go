package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/go/classroom-8/students/workshop-6/pkg/kafka"
)

const (
	topicName = "example-topic"
)

var brokers = []string{
	"127.0.0.1:9091",
	"127.0.0.1:9092",
	"127.0.0.1:9093",
}

func main() {
	producer, err := kafka.NewSyncProducer(brokers,
		kafka.WithIdempotent(),
		kafka.WithRequiredAcks(sarama.WaitForAll),
		kafka.WithProducerPartitioner(sarama.NewHashPartitioner),
		kafka.WithMaxOpenRequests(1),
		kafka.WithMaxRetries(5),
		kafka.WithRetryBackoff(10*time.Millisecond),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer producer.Close() // Не забываем освобождать ресурсы :)

	type Message struct {
		ID    int
		Value string
	}

	m := Message{
		ID:    1,
		Value: "example",
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	msg, err := kafka.BuildMessage(topicName, "key2", bytes, "x-header-example", "example-header-value")
	if err != nil {
		log.Fatal(err)
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("partition: %d, offset: %d", partition, offset)
}
