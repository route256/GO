package worker

import (
	"context"
	"log"
	"time"
)

type CurrencyExchangeRateUpdater interface {
	UpdateCurrencyExchangeRatesOn(ctx context.Context, time time.Time) error
}

type ConfigGetter interface {
	GetFrequencyExchangeRate() time.Duration
}

type CurrencyExchangeRateWorker struct {
	updater CurrencyExchangeRateUpdater
	config  ConfigGetter
}

func NewCurrencyExchangeRateWorker(updater CurrencyExchangeRateUpdater, config ConfigGetter) *CurrencyExchangeRateWorker {
	return &CurrencyExchangeRateWorker{updater: updater, config: config}
}

func (worker *CurrencyExchangeRateWorker) Run(ctx context.Context) {
	ticker := time.NewTicker(worker.config.GetFrequencyExchangeRate())

	go func() {
		for {
			select {
			// По заданию нужно отменять операцию на основании контекста
			case <-ctx.Done():
				log.Println("stopped receiving exchange rates")
				return
			case <-ticker.C:
				select {
				// Дополнительная проверка, что пока мы ждали таймер у нас не завершилась работать
				case <-ctx.Done():
					log.Println("stopped receiving exchange rates")
					return
				default:
					// Запрашиваем на текущую дату каждые N секунду
					if err := worker.updater.UpdateCurrencyExchangeRatesOn(ctx, time.Now()); err != nil {
						log.Println(err)
					}
				}
			}
		}
	}()
}
