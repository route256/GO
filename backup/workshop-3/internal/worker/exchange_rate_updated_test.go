package worker

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/worker/mocks"
)

func TestCurrencyExchangeRateWorkerRun(t *testing.T) {
	t.Run("after interrupting context receiving for current exchange rates will stop", func(t *testing.T) {
		t.Parallel()

		rateUpdateMock := mocks.NewMockCurrencyExchangeRateUpdater(gomock.NewController(t))
		configGetter := mocks.NewMockConfigGetter(gomock.NewController(t))
		configGetter.EXPECT().GetFrequencyExchangeRate().Return(30 * time.Millisecond)

		rateUpdateMock.EXPECT().
			UpdateCurrencyExchangeRatesOn(gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, date time.Time) error {
				assert.NoError(t, ctx.Err())
				return nil
			}).
			Times(4)

		worker := NewCurrencyExchangeRateWorker(rateUpdateMock, configGetter)

		ctx, cancel := context.WithTimeout(context.TODO(), 140*time.Millisecond)
		defer cancel()
		worker.Run(ctx)

		// За время дополнительно ожидания в случае кривого теста кол-во вызовов мока было бы больше 4
		time.Sleep(160 * time.Millisecond)
		assert.Error(t, ctx.Err())
	})
}
