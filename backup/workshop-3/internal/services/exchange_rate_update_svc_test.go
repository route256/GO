package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/domain"
	utils "gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/helpers/date"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/services/mocks"
)

// Лично я придерживаюсь go way нейминга тестов, а уже сам набор тестов разделять t.Run или табличными тестами:
// - тесты читаются проще так как в name осмысленный текст
func TestUpdateCurrencyExchangeRatesOn(t *testing.T) {
	t.Skip()
	t.Run("returns error if gateway CBR could not get exchange rate on the date", func(t *testing.T) {
		// Будем приучать студентов к запуску тестов в параллели ? :)
		t.Parallel()

		// ARRANGE (P.S. на работе я прям так секции редко делю)
		gatewayMock := mocks.NewMockExchangeRateFetcherGateway(gomock.NewController(t))
		svc := NewExchangeRateUpdateSvc(gatewayMock, nil, nil)

		date := utils.GetDate(time.Now())

		gatewayMock.EXPECT().
			FetchRatesOn(gomock.Any(), gomock.Eq(date)).
			Return(nil, errors.New("failed to fetch rates on current date"))

		// ACT
		err := svc.UpdateCurrencyExchangeRatesOn(context.TODO(), date)

		// ASSERT
		assert.EqualError(t, err, "failed to fetch rates on current date")
	})

	t.Run("we process only those currencies that are configured", func(t *testing.T) {
		t.Parallel()

		// ARRANGE
		gatewayMock := mocks.NewMockExchangeRateFetcherGateway(gomock.NewController(t))
		configMock := mocks.NewMockConfigGetter(gomock.NewController(t))
		storageMock := mocks.NewMockRateStorage(gomock.NewController(t))
		svc := NewExchangeRateUpdateSvc(gatewayMock, storageMock, configMock)

		date := utils.GetDate(time.Now())

		configMock.EXPECT().SupportedCurrencyCodes().Return([]string{"RUB", "EUR"})
		storageMock.EXPECT().AddRate(gomock.Any(), gomock.Eq(date), gomock.Eq(domain.Rate{Code: "RUB", Original: "0"})).Return(nil)
		storageMock.EXPECT().AddRate(gomock.Any(), gomock.Eq(date), gomock.Eq(domain.Rate{Code: "EUR", Original: "0"})).Return(nil)
		storageMock.EXPECT().AddRate(gomock.Any(), gomock.Eq(date), gomock.Eq(domain.Rate{Code: "USD", Original: "0"})).Times(0)

		gatewayMock.EXPECT().
			FetchRatesOn(gomock.Any(), gomock.Eq(date)).
			Return([]domain.Rate{{Code: "RUB", Original: "0"}, {Code: "USD", Original: "0"}, {Code: "EUR", Original: "0"}}, nil)

		// ACT
		err := svc.UpdateCurrencyExchangeRatesOn(context.TODO(), date)

		// ASSERT
		assert.NoError(t, err)
	})

	t.Run("penny converter should not block the saving of valid values", func(t *testing.T) {
		t.Parallel()

		// ARRANGE
		gatewayMock := mocks.NewMockExchangeRateFetcherGateway(gomock.NewController(t))
		configMock := mocks.NewMockConfigGetter(gomock.NewController(t))
		storageMock := mocks.NewMockRateStorage(gomock.NewController(t))
		svc := NewExchangeRateUpdateSvc(gatewayMock, storageMock, configMock)

		date := utils.GetDate(time.Now())

		configMock.EXPECT().SupportedCurrencyCodes().Return([]string{"RUB", "USD", "EUR"})
		storageMock.EXPECT().AddRate(gomock.Any(), gomock.Eq(date), gomock.Eq(domain.Rate{Code: "RUB", Original: "0"})).Return(nil)
		storageMock.EXPECT().AddRate(gomock.Any(), gomock.Eq(date), gomock.Eq(domain.Rate{Code: "EUR", Original: "0"})).Return(nil)
		storageMock.EXPECT().AddRate(gomock.Any(), gomock.Eq(date), gomock.Eq(domain.Rate{Code: "USD", Original: "0"})).Times(0)

		gatewayMock.EXPECT().
			FetchRatesOn(gomock.Any(), gomock.Eq(date)).
			Return([]domain.Rate{{Code: "RUB", Original: "0"}, {Code: "USD", Original: "value"}, {Code: "EUR", Original: "0"}}, nil)

		// ACT
		err := svc.UpdateCurrencyExchangeRatesOn(context.TODO(), date)

		// ASSERT
		assert.NoError(t, err)
	})

	t.Run("errors in saving currency rates should not block saving as a whole", func(t *testing.T) {
		t.Parallel()

		// ARRANGE
		gatewayMock := mocks.NewMockExchangeRateFetcherGateway(gomock.NewController(t))
		configMock := mocks.NewMockConfigGetter(gomock.NewController(t))
		storageMock := mocks.NewMockRateStorage(gomock.NewController(t))
		svc := NewExchangeRateUpdateSvc(gatewayMock, storageMock, configMock)

		date := utils.GetDate(time.Now())

		configMock.EXPECT().SupportedCurrencyCodes().Return([]string{"RUB", "USD", "EUR"})
		storageMock.EXPECT().AddRate(gomock.Any(), gomock.Eq(date), gomock.Eq(domain.Rate{Code: "RUB", Original: "0"})).Return(errors.New("save failed"))
		storageMock.EXPECT().AddRate(gomock.Any(), gomock.Eq(date), gomock.Eq(domain.Rate{Code: "EUR", Original: "0"})).Return(nil)
		storageMock.EXPECT().AddRate(gomock.Any(), gomock.Eq(date), gomock.Eq(domain.Rate{Code: "USD", Original: "0"})).Return(nil)

		gatewayMock.EXPECT().
			FetchRatesOn(gomock.Any(), gomock.Eq(date)).
			Return([]domain.Rate{{Code: "RUB", Original: "0"}, {Code: "USD", Original: "0"}, {Code: "EUR", Original: "0"}}, nil)

		// ACT
		err := svc.UpdateCurrencyExchangeRatesOn(context.TODO(), date)

		// ASSERT
		assert.NoError(t, err)
	})
}
