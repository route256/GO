package services

import (
	"context"
	"log"
	"time"

	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/domain"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/helpers/money"
)

// В нейминге не завязываемся на источник данные, это может быть как ЦБ РФ так и совершенно любое публичное API
type ExchangeRateFetcherGateway interface {
	FetchRatesOn(ctx context.Context, date time.Time) ([]domain.Rate, error)
}

type RateStorage interface {
	AddRate(ctx context.Context, date time.Time, rate domain.Rate) error
}

type ConfigGetter interface {
	SupportedCurrencyCodes() []string
}

type ExchangeRateUpdateSvc struct {
	gateway ExchangeRateFetcherGateway
	storage RateStorage
	config  ConfigGetter
}

func NewExchangeRateUpdateSvc(gateway ExchangeRateFetcherGateway, storage RateStorage, config ConfigGetter) *ExchangeRateUpdateSvc {
	return &ExchangeRateUpdateSvc{gateway: gateway, storage: storage, config: config}
}

// Вся логика лежит в отдельном сервисе по двум причинам, не хочется тащить в воркер доп логику и второе сервис
// потруебтся если пользователь будет запрашивать список рассходов на даты по которым у нас нету катировок
func (svc *ExchangeRateUpdateSvc) UpdateCurrencyExchangeRatesOn(ctx context.Context, time time.Time) error {
	rates, err := svc.gateway.FetchRatesOn(ctx, time)
	if err != nil {
		return err
	}

	//  Мы не хотим захламлять базу всем валютами, которые существуют
	supportedCurrencyCodes := svc.config.SupportedCurrencyCodes()
	//  Вот тут можно будет посмотреть как студенты сделают, ведь валют может быть больше 4, всего в мире ~149 валют + сколько криптовалют, которые могут поддерживаться ботом
	supportedCurrencyCodesAsMap := make(map[string]string, len(supportedCurrencyCodes))
	// Интересно кто-то из студентов решится использоваться https://github.com/samber/lo
	for _, code := range supportedCurrencyCodes {
		supportedCurrencyCodesAsMap[code] = code
	}

	for _, rate := range rates {
		if _, ok := supportedCurrencyCodesAsMap[rate.Code]; !ok {
			continue
		}

		// Курсы валют тоже храним в копейках, как и транзакции для более удобного использования
		// Но тут просядет точность 54.8827 -> 5488
		rate.Kopecks, err = money.ConvertStringAmountToKopecks(rate.Original)
		if err != nil {
			// ошибка не крит и вероятность ее тут словить мала, все таки мы должны доверять ЦБ РФ
			log.Println(err)
			continue
		}
		rate.Ts = time

		// Не тащим dto ЦБ РФ через все слои, работаем с доменом
		err = svc.storage.AddRate(ctx, time, rate)
		if err != nil {
			// ошибка не крит, мы можем перезапросить список через N часов, не стоит фейлить все корректировки
			log.Println(err)
		}
	}

	return nil
}
