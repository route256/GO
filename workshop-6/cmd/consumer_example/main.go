package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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
	consumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close() // Не забываем освобождать ресурсы :)

	consumer.ConsumeTopic(topicName, func(msg *sarama.ConsumerMessage) {
		// Your logic here
		fmt.Println("Read Topic: ", msg.Topic, " Partition: ", msg.Partition, " Offset: ", msg.Offset)
		fmt.Println("Received Key: ", string(msg.Key), " Value: ", string(msg.Value))
	})

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
}
