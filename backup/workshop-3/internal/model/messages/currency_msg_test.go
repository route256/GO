package messages

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	mock_messages "gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/mocks/messages"
)

func TestSetCurrency(t *testing.T) {
	expenseDB := mock_messages.NewMockExpenseDB(gomock.NewController(t))
	userDB := mock_messages.NewMockUserDB(gomock.NewController(t))
	currencySettings := mock_messages.NewMockConfigGetter(gomock.NewController(t))
	messageSender := mock_messages.NewMockMessageSender(gomock.NewController(t))
	model := New(messageSender, currencySettings, expenseDB, userDB, nil, nil)

	t.Run("return error if currency code unsupported", func(t *testing.T) {
		currencySettings.EXPECT().SupportedCurrencyCodes().Return([]string{"RUB", "EUR"}).Times(2)
		userDB.EXPECT().UserExist(gomock.Any(), int64(100)).Return(false)

		text, err := model.setCurrency(context.TODO(), Message{UserID: 100, Text: "/set_currency USD"})
		assert.NoError(t, err)
		assert.EqualValues(t, text, "Валюта USD не поддерживается, отправьте команду /set_currency с одним из значений [RUB EUR]")
	})

	t.Run("returns an error if an error occurred at the time of saving", func(t *testing.T) {
		currencySettings.EXPECT().SupportedCurrencyCodes().Return([]string{"RUB", "EUR"})
		userDB.EXPECT().UserExist(gomock.Any(), int64(100)).Return(true)
		userDB.EXPECT().ChangeDefaultCurrency(gomock.Any(), int64(100), "EUR").Return(errors.New("error"))
		_, err := model.setCurrency(context.TODO(), Message{UserID: 100, Text: "/set_currency EUR"})

		assert.EqualError(t, err, "failed to change user currency")
	})

	t.Run("a hint on how to use the bot will be returned for new users", func(t *testing.T) {
		currencySettings.EXPECT().SupportedCurrencyCodes().Return([]string{"RUB", "EUR"})
		userDB.EXPECT().UserExist(gomock.Any(), int64(100)).Return(false)
		userDB.EXPECT().ChangeDefaultCurrency(gomock.Any(), int64(100), "EUR").Return(nil)
		text, err := model.setCurrency(context.TODO(), Message{UserID: 100, Text: "/set_currency EUR"})

		assert.NoError(t, err)
		assert.EqualValues(t, helpMessage, text)
	})

	t.Run("hint on how to use the bot will be returned for new users", func(t *testing.T) {
		currencySettings.EXPECT().SupportedCurrencyCodes().Return([]string{"RUB", "EUR"})
		userDB.EXPECT().UserExist(gomock.Any(), int64(100)).Return(false)
		userDB.EXPECT().ChangeDefaultCurrency(gomock.Any(), int64(100), "EUR").Return(nil)
		text, err := model.setCurrency(context.TODO(), Message{UserID: 100, Text: "/set_currency EUR"})

		assert.NoError(t, err)
		assert.EqualValues(t, helpMessage, text)
	})

	t.Run("notification about change currency will be returned for old  users", func(t *testing.T) {
		currencySettings.EXPECT().SupportedCurrencyCodes().Return([]string{"RUB", "EUR"})
		userDB.EXPECT().UserExist(gomock.Any(), int64(100)).Return(true)
		userDB.EXPECT().ChangeDefaultCurrency(gomock.Any(), int64(100), "EUR").Return(nil)
		text, err := model.setCurrency(context.TODO(), Message{UserID: 100, Text: "/set_currency EUR"})

		assert.NoError(t, err)
		assert.EqualValues(t, "Установлена валюта по умолчанию EUR", text)
	})
}

func TestChangeDefaultCurrency(t *testing.T) {
	expenseDB := mock_messages.NewMockExpenseDB(gomock.NewController(t))
	userDB := mock_messages.NewMockUserDB(gomock.NewController(t))
	currencySettings := mock_messages.NewMockConfigGetter(gomock.NewController(t))
	messageSender := mock_messages.NewMockMessageSender(gomock.NewController(t))
	model := New(messageSender, currencySettings, expenseDB, userDB, nil, nil)

	t.Run("return command for changing  default currency", func(t *testing.T) {
		currencySettings.EXPECT().SupportedCurrencyCodes().Return([]string{"RUB", "USD", "EUR"})
		text, buttons := model.changeDefaultCurrency()

		exceptedCurrencies := []map[string]string{
			{
				"RUB": "/set_currency RUB",
				"USD": "/set_currency USD",
				"EUR": "/set_currency EUR",
			},
		}

		assert.EqualValues(t, "Выберите валюту в которой будете производить расходы", text)
		assert.EqualValues(t, exceptedCurrencies, buttons)
	})
}
