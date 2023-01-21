package messages

import (
	"context"
	"fmt"
	"log"
	"strings"
)

func (s *Model) setCurrency(ctx context.Context, msg Message) (text string, err error) {
	currencyCode := strings.Trim(strings.TrimPrefix(msg.Text, "/set_currency"), " ")
	userExist := s.userDB.UserExist(ctx, msg.UserID)

	supportedCurrencyCodesAsMap := make(map[string]any)
	for _, supportedCurrencyCode := range s.config.SupportedCurrencyCodes() {
		supportedCurrencyCodesAsMap[supportedCurrencyCode] = struct{}{}
	}

	if _, ok := supportedCurrencyCodesAsMap[currencyCode]; ok {
		if err = s.userDB.ChangeDefaultCurrency(ctx, msg.UserID, currencyCode); err != nil {
			log.Println(err)
			return "", ErrImpossibleToChangeUserCurrency
		}

		if userExist {
			return fmt.Sprintf("Установлена валюта по умолчанию %s", currencyCode), nil
		} else {
			return helpMessage, nil
		}

	}

	return fmt.Sprintf("Валюта %s не поддерживается, отправьте команду /set_currency с одним из значений %v", currencyCode, s.config.SupportedCurrencyCodes()), nil
}

func (s *Model) changeDefaultCurrency() (text string, buttons []map[string]string) {
	currencyCodes := s.config.SupportedCurrencyCodes()

	rows := make(map[string]string, len(currencyCodes))
	for _, currencyCode := range currencyCodes {
		rows[currencyCode] = fmt.Sprintf("/set_currency %s", currencyCode)
	}

	return "Выберите валюту в которой будете производить расходы", []map[string]string{rows}
}
