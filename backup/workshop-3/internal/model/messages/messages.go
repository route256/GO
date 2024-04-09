package messages

import (
	"context"
	"errors"
	"time"

	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/domain"
)

const (
	dateFormat = "02.01.2006"
)

const (
	InvalidCommandMessage       = "Неверный формат команды, исправьте и повторите команду"
	InvalidAmountMessage        = "Неверный формат суммы, исправьте и повторите команду"
	InvalidDateMessage          = "Неверный формат даты, исправьте и повторите команду"
	FailedWriteMessage          = "Не удалось записать расход, повторите попытку позже"
	FailedMessage               = "Я временно не работаю, повторите попытку позже"
	FailedChangeCurrencyMessage = "Не удалось изменить текущую валюту, повторите попытку позже"
)

var (
	ErrInvalidCommand                 = errors.New("invalid command")
	ErrInvalidAmount                  = errors.New("invalid amount")
	ErrInvalidDate                    = errors.New("invalid date")
	ErrWriteToDatabase                = errors.New("failed to write to the database")
	ErrGetRecordsInDatabase           = errors.New("failed to get records from the database")
	ErrImpossibleToChangeUserCurrency = errors.New("failed to change user currency")
)

type CurrencyExchangeRateUpdater interface {
	UpdateCurrencyExchangeRatesOn(ctx context.Context, time time.Time) error
}

type MessageSender interface {
	SendMessage(text string, userID int64, buttons ...map[string]string) error
}

type ExpenseDB interface {
	AddExpense(ctx context.Context, userID int64, kopecks int64, title string, date time.Time) error
	GetExpenses(ctx context.Context, userID int64) ([]domain.Expense, error)
}

type UserDB interface {
	UserExist(ctx context.Context, userID int64) bool
	ChangeDefaultCurrency(ctx context.Context, userID int64, currency string) error
	GetDefaultCurrency(ctx context.Context, userID int64) (string, error)
}

type RateDB interface {
	GetRate(ctx context.Context, code string, date time.Time) (*domain.Rate, error)
}

type ConfigGetter interface {
	SupportedCurrencyCodes() []string
	GetBaseCurrency() string
}

type Model struct {
	tgClient    MessageSender
	expenseDB   ExpenseDB
	userDB      UserDB
	rateDB      RateDB
	config      ConfigGetter
	rateUpdater CurrencyExchangeRateUpdater
}

func New(tgClient MessageSender, config ConfigGetter, expenseDB ExpenseDB, userDB UserDB, rateDB RateDB, rateUpdater CurrencyExchangeRateUpdater) *Model {
	return &Model{
		tgClient:    tgClient,
		config:      config,
		expenseDB:   expenseDB,
		userDB:      userDB,
		rateDB:      rateDB,
		rateUpdater: rateUpdater,
	}
}

type Message struct {
	Text   string
	UserID int64
}
