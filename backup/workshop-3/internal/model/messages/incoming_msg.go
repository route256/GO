package messages

import (
	"context"
	"errors"
	"strings"
)

const newUser = `Привет! Я буду помогать ввести твою бухгалтерию. Но перед началом работы тебе нужны выбрать валюту по умолчанию в которой ты производишь расходы`

const introMessage = "Привет! Я умею учитывать твои траты.\n\n" + helpMessage

const helpMessage = `Для работы с ботом тебе могут потребоваться следующие команды:
Чтобы изменить выбранную валюту необходимо выполнить команду /change_currency

Чтобы добавить новый расход, отправь мне сообщение в формате:
/add цена; описание; дата (дд.мм.гггг, опционально) - если не указать дату, то расход будет добавлен на сегодняшний день

Чтобы посмотреть расходы отправь:
/list -  за всё время.
/list_day - за день.
/list_week - за неделю.
/list_month - за месяц.
/list_year - за год.`

const unknownMessage = `Неизвестная команда. Чтобы посмотреть список команд отправь /help`

func (s *Model) IncomingMessage(ctx context.Context, msg Message) error {
	if !s.userDB.UserExist(ctx, msg.UserID) && !strings.HasPrefix(msg.Text, "/set_currency") {
		_, buttons := s.changeDefaultCurrency()

		return s.tgClient.SendMessage(newUser, msg.UserID, buttons...)
	}

	if msg.Text == "/start" {
		return s.tgClient.SendMessage(introMessage, msg.UserID)
	}

	if msg.Text == "/help" {
		return s.tgClient.SendMessage(helpMessage, msg.UserID)
	}

	if strings.HasPrefix(msg.Text, "/add") {
		answer, err := s.addExpense(ctx, msg)
		if err == nil {
			return s.tgClient.SendMessage(answer, msg.UserID)
		}

		if errors.Is(err, ErrInvalidCommand) {
			return s.tgClient.SendMessage(InvalidCommandMessage, msg.UserID)
		}

		if errors.Is(err, ErrInvalidAmount) {
			return s.tgClient.SendMessage(InvalidAmountMessage, msg.UserID)
		}

		if errors.Is(err, ErrInvalidDate) {
			return s.tgClient.SendMessage(InvalidDateMessage, msg.UserID)
		}

		if errors.Is(err, ErrWriteToDatabase) {
			return s.tgClient.SendMessage(FailedWriteMessage, msg.UserID)
		}

		// fallback error
		return s.tgClient.SendMessage(FailedMessage, msg.UserID)
	}

	if strings.HasPrefix(msg.Text, "/list") {
		answer, err := s.listExpenses(ctx, msg)
		if err == nil {
			return s.tgClient.SendMessage(answer, msg.UserID)
		}

		if errors.Is(err, ErrGetRecordsInDatabase) {
			return s.tgClient.SendMessage(FailedGetListExpensesMessage, msg.UserID)
		}

		// fallback error
		return s.tgClient.SendMessage(FailedMessage, msg.UserID)
	}

	if strings.HasPrefix(msg.Text, "/set_currency") {
		answer, err := s.setCurrency(ctx, msg)
		if err == nil {
			return s.tgClient.SendMessage(answer, msg.UserID)
		}

		if errors.Is(err, ErrImpossibleToChangeUserCurrency) {
			return s.tgClient.SendMessage(FailedChangeCurrencyMessage, msg.UserID)
		}

		// fallback error
		return s.tgClient.SendMessage(FailedMessage, msg.UserID)
	}

	if strings.HasPrefix(msg.Text, "/change_currency") {
		answer, buttons := s.changeDefaultCurrency()
		return s.tgClient.SendMessage(answer, msg.UserID, buttons...)
	}

	return s.tgClient.SendMessage(unknownMessage, msg.UserID)
}
