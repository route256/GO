package usecase

import (
	gobytes "bytes"
	"context"
	"strconv"
	"strings"
	"workshopcache/config"
	"workshopcache/domain/entity"
	"workshopcache/gateway"

	"go.uber.org/zap"
)

type CryptoUseCase struct {
	cryptoGw gateway.CryptoGw
	cacheGw  gateway.CacheGW
	rule     int64
}

func NewCryptoUseCase(cfg *config.CryptoConfig, cryptoGw gateway.CryptoGw, cacheGw gateway.CacheGW) *CryptoUseCase {
	var rule int64
	if cfg != nil {
		rule = cfg.Rule
	}
	return &CryptoUseCase{
		cryptoGw: cryptoGw,
		cacheGw:  cacheGw,
		rule:     rule,
	}
}

func (uc *CryptoUseCase) Mine(ctx context.Context, data string, logger *zap.Logger) (int64, error) {

	var PoW int64
	var hash string

	val, err := uc.cacheGw.Get(ctx, data)
	if val == "" {
		logger.Debug("No value in cache for the given key", zap.String("key", data))

		var req = entity.MineRequest{
			Bytes: []byte(data),
		}

		for {
			PoW += 1
			finalString := gobytes.Join([][]byte{req.Bytes, []byte(strconv.FormatInt(PoW, 10))}, []byte("-"))
			hash = uc.cryptoGw.GenerateHash(finalString)
			if strings.HasPrefix(hash, strconv.FormatInt(uc.rule, 10)) {

				var PoWStr = strconv.FormatInt(PoW, 10)
				logger.Debug("Add data to the cache", zap.String("key", data), zap.String("value", PoWStr))
				err := uc.cacheGw.Set(ctx, data, PoWStr)
				if err != nil {
					return 0, err
				}
				return PoW, nil

			}
		}
	}
	logger.Debug("Found value in cache for the given key!", zap.String("key", data))
	PoW, err = strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return PoW, err

}
