package main

import (
	"log"

	"gitlab.ozon.dev/classroom-4/teachers/workshop-1-bot/internal/clients/tg"
	"gitlab.ozon.dev/classroom-4/teachers/workshop-1-bot/internal/config"
	"gitlab.ozon.dev/classroom-4/teachers/workshop-1-bot/internal/model/messages"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	tgClient, err := tg.New(config)
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel)
}
