package main

import (
	"context"
	"workshopcache/config"
	"workshopcache/domain/usecase"
	"workshopcache/gateway"
	"workshopcache/infrastructure/cache"
	"workshopcache/infrastructure/log"
	"workshopcache/infrastructure/net"

	"go.uber.org/zap"
)

func main() {

	cfg := config.InitConfig()

	appLogger, err := log.NewLogger(cfg.Log.Level)
	if err != nil {
		panic(err)
	}

	appLogger.Info("The app is starting...")
	fatalErrCh := make(chan error, 2)

	redisClient := cache.NewCacheClient(cfg.Cache)
	ctx := context.Background()

	cryptoGw := gateway.NewCryptoGateway()
	cacheGw := gateway.NewCacheGateway(redisClient)
	cryptoUc := usecase.NewCryptoUseCase(cfg.Crypto, cryptoGw, cacheGw)

	go func() {

		err := net.NewHTTPConnection(ctx, cryptoUc, cfg, appLogger, fatalErrCh)
		if err != nil {
			appLogger.Fatal("The app could not established the http connection. Exiting")
		}
		appLogger.Info("The app has started")

	}()

	select {
	case err := <-fatalErrCh:
		appLogger.Error("Received error from functional unit", zap.Error(err))
	}
}
