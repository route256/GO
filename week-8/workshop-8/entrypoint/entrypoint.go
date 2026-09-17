package entrypoint

import (
	"context"
	"net/http"
	"strconv"
	"workshopcache/domain/usecase"

	"go.uber.org/zap"
)

func Mine(ctx context.Context, cryptoUc *usecase.CryptoUseCase, logger *zap.Logger) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		data := r.URL.Query().Get("data")
		logger.Debug("Hit the /Mine endpoint", zap.String("data", data))

		PoW, err := cryptoUc.Mine(ctx, data, logger)
		if err != nil {
			logger.Error(err.Error())
			return
		}
		logger.Debug("Calculated PoW for the given data", zap.Int64("pow", PoW))

		response := []byte(strconv.FormatInt(PoW, 10))
		_, err = w.Write(response)
		if err != nil {
			logger.Error(err.Error())
			return
		}
	}
	return http.HandlerFunc(fn)
}
