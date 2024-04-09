package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/config"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/database"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/infrastructure/cbr_gateway"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/infrastructure/tg_gateway"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/model/messages"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/services"
	worker "gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/worker"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=pass sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// DATABASE
	expenseDB, err := database.NewExpenseDB()
	if err != nil {
		log.Fatal("database init failed:", err)
	}
	userDB, err := database.NewUserDB()
	if err != nil {
		log.Fatal("database init failed:", err)
	}
	rateDB := database.NewRateDB(db)

	// GATEWAY

	cbrRateAPIGateway := cbr_gateway.New()
	tgAPIGateway, err := tg_gateway.New(config)
	if err != nil {
		log.Fatal("tg bot create failed:", err)
	}

	// SERVICES

	exchangeRateUpdateSvc := services.NewExchangeRateUpdateSvc(cbrRateAPIGateway, rateDB, config)

	// COMMANDS
	messageProcessor := messages.New(tgAPIGateway, config, expenseDB, userDB, rateDB, exchangeRateUpdateSvc)

	// WORKERS
	currencyExchangeRateWorker := worker.NewCurrencyExchangeRateWorker(exchangeRateUpdateSvc, config)
	messageListenerWorker := worker.NewMessageListenerWorker(tgAPIGateway, messageProcessor)

	currencyExchangeRateWorker.Run(ctx)
	messageListenerWorker.Run(ctx)
}
