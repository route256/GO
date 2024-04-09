package net

import (
	"context"
	"fmt"
	"net/http"
	"workshopcache/config"
	"workshopcache/domain/usecase"
	"workshopcache/entrypoint"

	"go.uber.org/zap"
)

const serviceName = "http"

func NewHTTPConnection(ctx context.Context, cryptoUc *usecase.CryptoUseCase, cfg *config.Config, logger *zap.Logger, fatalErrCh chan error) error {

	serviceLogger := logger.Named(serviceName)
	serviceLogger.Info("Establishing the http connection ...")

	mux := http.NewServeMux()
	mineHandler := entrypoint.Mine(ctx, cryptoUc, logger)
	mux.Handle("/mine", mineHandler)

	var addr string
	if cfg != nil && cfg.App != nil {
		addr = fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)
	}

	go func() {
		err := http.ListenAndServe(addr, mux)
		if err != nil {
			fatalErrCh <- err
		}
	}()

	serviceLogger.Info("Established the http connection successfully.")

	return nil

}
