package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Наш обработчик реализующий интерфейс sarama.ConsumerGroupHandler
	consumerGroupHandler := kafka.NewConsumerGroupHandler( /**/ )

	// Создаем коньюмер группу
	consumerGroup, err := kafka.NewConsumerGroup(
		brokers,
		"consumer-group-example-1",
		[]string{topicName},
		consumerGroupHandler,
	)
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		// запускаем вычитку сообщений
		consumerGroup.Run(ctx)
	}()

	<-consumerGroupHandler.Ready() // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	var (
		consumptionIsPaused = false
		keepRunning         = true
	)
	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(consumerGroup, &consumptionIsPaused)
		}
	}

	cancel()
	wg.Wait()

	if err = consumerGroup.Close(); err != nil {
		log.Fatalf("Error closing consumer group: %v", err)
	}
}

func toggleConsumptionFlow(cg sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		cg.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		cg.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}
