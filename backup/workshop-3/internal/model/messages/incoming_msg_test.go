package messages

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/domain"
	mocks_msg "gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/mocks/messages"
)

func Test_OnAddCommand_ShouldAnswerWithAddedMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := minimock.NewController(t)
	defer m.Finish()

	userID := int64(123)

	testCases := []struct {
		name    string
		command string
		amount  string
		kopecks int64
		title   string
		date    interface{}
		answer  string
	}{
		{
			name:    "normal",
			amount:  "100.0",
			kopecks: 10000,
			title:   "расход",
			date:    parseDate("01.10.2022"),
			command: "/add 100.0; расход; 01.10.2022",
			answer:  "Расход добавлен: 100.00 RUB расход 01.10.2022",
		},
		{
			name:    "without title",
			amount:  "100.0",
			kopecks: 10000,
			title:   "",
			date:    parseDate("01.10.2022"),
			command: "/add 100.0; ;01.10.2022",
			answer:  "Расход добавлен: 100.00 RUB  01.10.2022",
		},
		{
			name:    "without date",
			amount:  "100.0",
			kopecks: 10000,
			title:   "расход",
			date:    gomock.Any(),
			command: "/add 100.0; расход;",
			answer:  "Расход добавлен: 100.00 RUB расход " + time.Now().UTC().Format(dateFormat),
		},
		{
			name:    "only amount with semicolon",
			amount:  "100.0",
			kopecks: 10000,
			title:   "",
			date:    gomock.Any(),
			command: "/add 100.0;",
			answer:  "Расход добавлен: 100.00 RUB  " + time.Now().UTC().Format(dateFormat),
		},
		{
			name:    "only amount without semicolon",
			amount:  "100.0",
			kopecks: 10000,
			title:   "",
			date:    gomock.Any(),
			command: "/add 100.0",
			answer:  "Расход добавлен: 100.00 RUB  " + time.Now().UTC().Format(dateFormat),
		},
		{
			name:    "without amount",
			amount:  "",
			kopecks: 0,
			title:   "расход",
			date:    gomock.Any(),
			command: "/add ; расход; 01.10.2022",
			answer:  InvalidAmountMessage,
		},
		{
			name:    "invalid amount",
			amount:  "100.0.0",
			kopecks: 0,
			title:   "расход",
			date:    gomock.Any(),
			command: "/add 100.0.0; расход; 01.10.2022",
			answer:  InvalidAmountMessage,
		},
		{
			name:    "invalid date",
			amount:  "100.0",
			kopecks: 10000,
			title:   "расход",
			date:    gomock.Any(),
			command: "/add 100.0; расход; 01.24.2022",
			answer:  InvalidDateMessage,
		},
		{
			name:    "invalid date format",
			amount:  "100.0",
			kopecks: 10000,
			title:   "расход",
			date:    gomock.Any(),
			command: "/add 100.0; расход; 01.10.2022.0",
			answer:  InvalidDateMessage,
		},
		{
			name:    "empty command",
			command: "/add",
			date:    gomock.Any(),
			answer:  InvalidCommandMessage,
		},
	}

	baseCurrency := "RUB"
	sender := mocks_msg.NewMockMessageSender(ctrl)
	config := mocks_msg.NewMockConfigGetter(ctrl)
	expenseDB := mocks_msg.NewMockExpenseDB(ctrl)
	userDB := mocks_msg.NewMockUserDB(ctrl)
	rateDB := NewRateDBMock(m)
	updater := mocks_msg.NewMockCurrencyExchangeRateUpdater(ctrl)
	model := New(sender, config, expenseDB, userDB, rateDB, updater)

	config.EXPECT().GetBaseCurrency().Return(baseCurrency).AnyTimes()
	userDB.EXPECT().GetDefaultCurrency(gomock.Any(), userID).Return(baseCurrency, nil).AnyTimes()
	userDB.EXPECT().UserExist(gomock.Any(), gomock.Any()).Return(true).AnyTimes()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expenseDB.EXPECT().AddExpense(gomock.Any(), userID, tc.kopecks, tc.title, tc.date).Return(nil).AnyTimes()
			sender.EXPECT().SendMessage(tc.answer, userID).Return(nil)

			err := model.IncomingMessage(context.TODO(), Message{
				Text:   tc.command,
				UserID: userID,
			})

			assert.NoError(t, err)
		})
	}
}

func Test_OnAddCommand_WithConverting(t *testing.T) {
	ctx := context.Background()
	m := minimock.NewController(t)
	defer m.Finish()

	sender := mocks_msg.NewMockMessageSender(gomock.NewController(t))
	config := mocks_msg.NewMockConfigGetter(gomock.NewController(t))
	expenseDB := mocks_msg.NewMockExpenseDB(gomock.NewController(t))
	userDB := mocks_msg.NewMockUserDB(gomock.NewController(t))
	rateDB := NewRateDBMock(m)
	updater := mocks_msg.NewMockCurrencyExchangeRateUpdater(gomock.NewController(t))
	model := New(sender, config, expenseDB, userDB, rateDB, updater)

	config.EXPECT().GetBaseCurrency().Return("RUB").Times(3)
	userDB.EXPECT().UserExist(gomock.Any(), gomock.Any()).Return(true).AnyTimes()
	userDB.EXPECT().GetDefaultCurrency(gomock.Any(), gomock.Any()).Return("USD", nil).Times(3)
	tcs := []struct {
		usd float64
		rub float64
	}{
		{usd: 1, rub: 60.16},
		{usd: 10, rub: 601.6},
		{usd: 15, rub: 902.4},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(fmt.Sprintf("%f -> %f", tc.usd, tc.rub), func(t *testing.T) {
			t.Parallel()

			rateDB.GetRateMock.Expect(ctx, "USD", parseDate("01.10.2022")).Return(
				&domain.Rate{
					Nominal: 1,
					Kopecks: 6016,
				}, nil)

			expenseDB.EXPECT().AddExpense(gomock.Any(), gomock.Any(), int64(tc.rub*100), "купил молока на всех", parseDate("01.10.2022")).Return(nil).AnyTimes()
			sender.EXPECT().SendMessage(fmt.Sprintf("Расход добавлен: %.2f USD купил молока на всех 01.10.2022", tc.usd), gomock.Any()).Return(nil)

			err := model.IncomingMessage(context.TODO(), Message{
				Text:   fmt.Sprintf("/add %.2f; купил молока на всех; 01.10.2022", tc.usd),
				UserID: 100,
			})

			assert.NoError(t, err)
		})
	}
}

func parseDate(date string) time.Time {
	v, _ := time.ParseInLocation(dateFormat, "01.10.2022", time.UTC)
	return v
}

func Test_OnUnknownCommand_ShouldAnswerWithUnknownMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sender := mocks_msg.NewMockMessageSender(ctrl)
	expenseDB := mocks_msg.NewMockExpenseDB(ctrl)
	userDB := mocks_msg.NewMockUserDB(gomock.NewController(t))

	sender.EXPECT().SendMessage(unknownMessage, int64(123))
	userDB.EXPECT().UserExist(gomock.Any(), gomock.Any()).Return(true).AnyTimes()

	model := New(sender, nil, expenseDB, userDB, nil, nil)

	err := model.IncomingMessage(context.TODO(), Message{
		Text:   "some text",
		UserID: 123,
	})

	assert.NoError(t, err)
}

func Test_OnListCommand_ShouldAnswerWithListMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := minimock.NewController(t)
	defer m.Finish()

	expenses := []domain.Expense{
		{
			Amount: 14.0,
			Title:  "потерял",
			Date:   time.Now().AddDate(-1, -2, 0).UTC(), // 1 год и 2 месяц назад
		}, {
			Amount: 5000.0,
			Title:  "залил бензин",
			Date:   time.Now().AddDate(0, -9, 0).UTC(), // 9 месяцев назад
		}, {
			Amount: 200.0,
			Title:  "сходил в кино",
			Date:   time.Now().AddDate(0, 0, -15).UTC(), // 15 дней назад
		}, {
			Amount: 1100.0,
			Title:  "сходил в магазин за продуктами",
			Date:   time.Now().AddDate(0, 0, -5).UTC(), // 5 дней назад
		}, {
			Amount: 136.0,
			Title:  "сходил за хлебом",
			Date:   time.Now().AddDate(0, 0, -1).UTC(), // вчера
		}, {
			Amount: 500.0,
			Title:  "поел",
			Date:   time.Now().UTC(), // сегодня
		},
	}

	baseCurrency := "RUB"
	sender := mocks_msg.NewMockMessageSender(ctrl)
	config := mocks_msg.NewMockConfigGetter(ctrl)
	expenseDB := mocks_msg.NewMockExpenseDB(ctrl)
	userDB := mocks_msg.NewMockUserDB(ctrl)
	rateDB := NewRateDBMock(m)
	updater := mocks_msg.NewMockCurrencyExchangeRateUpdater(ctrl)
	model := New(sender, config, expenseDB, userDB, rateDB, updater)

	config.EXPECT().GetBaseCurrency().Return(baseCurrency).AnyTimes()
	userDB.EXPECT().GetDefaultCurrency(gomock.Any(), gomock.Any()).Return(baseCurrency, nil).AnyTimes()
	userDB.EXPECT().UserExist(gomock.Any(), gomock.Any()).Return(true).AnyTimes()

	testCases := []struct {
		name     string
		command  string
		interval string
		include  []int64
	}{
		{
			name:     "За все время",
			command:  "/list",
			interval: "всё время",
			include:  []int64{0, 1, 2, 3, 4, 5},
		}, {
			name:     "За год",
			command:  "/list_year",
			interval: "год",
			include:  []int64{1, 2, 3, 4, 5},
		}, {
			name:     "За месяц",
			command:  "/list_month",
			interval: "месяц",
			include:  []int64{2, 3, 4, 5},
		}, {
			name:     "За неделю",
			command:  "/list_week",
			interval: "неделю",
			include:  []int64{3, 4, 5},
		}, {
			name:     "За день",
			command:  "/list_day",
			interval: "день",
			include:  []int64{5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			list := ""
			total := int64(0)
			for _, v := range tc.include {
				total += expenses[v].Amount
				list += fmt.Sprintf(formatItemListMessage, fmt.Sprintf("%.2f", float64(expenses[v].Amount)/100), baseCurrency, expenses[v].Title, expenses[v].Date.Format(dateFormat))
			}

			msg := fmt.Sprintf(formatTotalListMessage, tc.interval, fmt.Sprintf("%.2f", float64(total)/100), baseCurrency) + list

			expenseDB.EXPECT().GetExpenses(gomock.Any(), int64(123)).Return(expenses, nil)
			sender.EXPECT().SendMessage(msg, int64(123)).Return(nil)

			err := model.IncomingMessage(context.TODO(), Message{
				Text:   tc.command,
				UserID: 123,
			})
			assert.NoError(t, err)
		})
	}

	t.Run("failed get records", func(t *testing.T) {
		expenseDB.EXPECT().GetExpenses(gomock.Any(), int64(123)).Return(nil, errors.New("some error"))
		sender.EXPECT().SendMessage(FailedGetListExpensesMessage, int64(123)).Return(nil)

		err := model.IncomingMessage(context.TODO(), Message{
			Text:   "/list",
			UserID: 123,
		})

		assert.NoError(t, err)
	})
}

func Test_OnListCommand_WithConverting(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := minimock.NewController(t)
	defer m.Finish()

	t.Run("base case", func(t *testing.T) {

		date := time.Now().UTC()

		expenses := []domain.Expense{
			{
				Amount: 14700,
				Title:  "сходил в магазин",
				Date:   date, // сегодня
			},
			{
				Amount: 31200,
				Title:  "Купил воды",
				Date:   date, // сегодня
			},
			{
				Amount: 6200,
				Title:  "поел",
				Date:   date, // сегодня
			},
		}

		sender := mocks_msg.NewMockMessageSender(ctrl)
		config := mocks_msg.NewMockConfigGetter(ctrl)
		expenseDB := mocks_msg.NewMockExpenseDB(ctrl)
		userDB := mocks_msg.NewMockUserDB(ctrl)
		rateDB := NewRateDBMock(m)
		updater := mocks_msg.NewMockCurrencyExchangeRateUpdater(ctrl)
		model := New(sender, config, expenseDB, userDB, rateDB, updater)

		expenseDB.EXPECT().GetExpenses(gomock.Any(), gomock.Any()).Return(expenses, nil)
		config.EXPECT().GetBaseCurrency().Return("RUB")
		userDB.EXPECT().UserExist(gomock.Any(), gomock.Any()).Return(true)
		userDB.EXPECT().GetDefaultCurrency(gomock.Any(), gomock.Any()).Return("USD", nil)
		rateDB.GetRateMock.Expect(ctx, "USD", date).
			Return(&domain.Rate{
				Nominal: 1,
				Kopecks: 6016,
			}, nil)

		msg := fmt.Sprintf("Расходов за день: 8.65 USD\n------------------------------\n- 2.44 USD сходил в магазин (%[1]s)\n- 5.18 USD Купил воды (%[1]s)\n- 1.03 USD поел (%[1]s)", date.Format("02.01.2006"))
		sender.EXPECT().SendMessage(gomock.Eq(msg), gomock.Any()).Return(nil)

		err := model.IncomingMessage(context.TODO(), Message{
			Text:   "/list_day",
			UserID: 100,
		})

		assert.NoError(t, err)
	})

	t.Run("rate with nominal", func(t *testing.T) {
		date := time.Now().UTC()
		expenses := []domain.Expense{
			{
				Amount: 14700,
				Title:  "сходил в магазин",
				Date:   date, // сегодня
			},
		}

		sender := mocks_msg.NewMockMessageSender(ctrl)
		config := mocks_msg.NewMockConfigGetter(ctrl)
		expenseDB := mocks_msg.NewMockExpenseDB(ctrl)
		userDB := mocks_msg.NewMockUserDB(ctrl)
		rateDB := NewRateDBMock(m)
		updater := mocks_msg.NewMockCurrencyExchangeRateUpdater(ctrl)
		model := New(sender, config, expenseDB, userDB, rateDB, updater)

		expenseDB.EXPECT().GetExpenses(gomock.Any(), gomock.Any()).Return(expenses, nil)
		config.EXPECT().GetBaseCurrency().Return("RUB")
		userDB.EXPECT().UserExist(gomock.Any(), gomock.Any()).Return(true)
		userDB.EXPECT().GetDefaultCurrency(gomock.Any(), gomock.Any()).Return("CNY", nil)
		rateDB.GetRateMock.Expect(ctx, "CNY", date).
			Return(&domain.Rate{
				Nominal:  10,
				Original: "87.0042",
				Kopecks:  8700,
			}, nil)

		msg := fmt.Sprintf("Расходов за день: 16.89 CNY\n------------------------------\n- 16.89 CNY сходил в магазин (%s)", date.Format("02.01.2006"))
		sender.EXPECT().SendMessage(gomock.Eq(msg), gomock.Any()).Return(nil)

		err := model.IncomingMessage(context.TODO(), Message{
			Text:   "/list_day",
			UserID: 100,
		})

		assert.NoError(t, err)
	})
}

func TestStartCommand(t *testing.T) {
	t.Run("if user starts using for first time then we have to show him message about choosing currency", func(t *testing.T) {
		t.Parallel()

		userDB := mocks_msg.NewMockUserDB(gomock.NewController(t))
		currencySettings := mocks_msg.NewMockConfigGetter(gomock.NewController(t))
		messageSender := mocks_msg.NewMockMessageSender(gomock.NewController(t))
		model := New(messageSender, currencySettings, nil, userDB, nil, nil)

		userDB.EXPECT().UserExist(gomock.Any(), int64(10)).Return(false)
		currencySettings.EXPECT().SupportedCurrencyCodes().Return([]string{"RUB", "USD", "EUR"})
		exceptedCurrencies := []map[string]string{
			{
				"RUB": "/set_currency RUB",
				"USD": "/set_currency USD",
				"EUR": "/set_currency EUR",
			},
		}
		messageSender.EXPECT().SendMessage(gomock.Eq(newUser), gomock.Any(), gomock.Eq(exceptedCurrencies)).Return(nil)
		err := model.IncomingMessage(context.TODO(), Message{UserID: 10, Text: "/start"})
		assert.NoError(t, err)
	})

	t.Run("if user already exist we show help message", func(t *testing.T) {
		t.Parallel()

		userDB := mocks_msg.NewMockUserDB(gomock.NewController(t))
		currencySettings := mocks_msg.NewMockConfigGetter(gomock.NewController(t))
		messageSender := mocks_msg.NewMockMessageSender(gomock.NewController(t))
		model := New(messageSender, currencySettings, nil, userDB, nil, nil)

		userDB.EXPECT().UserExist(gomock.Any(), int64(10)).Return(true)

		messageSender.EXPECT().SendMessage(gomock.Eq(introMessage), gomock.Any()).Return(nil)
		err := model.IncomingMessage(context.TODO(), Message{UserID: 10, Text: "/start"})
		assert.NoError(t, err)
	})
}
