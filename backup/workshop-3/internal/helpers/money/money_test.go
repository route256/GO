package money_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/helpers/money"
)

// func main() {
// 	amount, err := money.ConvertStringAmountToKopecks("1,000,500.100")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(amount)
// 	fmt.Println(money.ConvertKopecksToAmount(amount))
// }

func TestConvert(t *testing.T) {

	testCases := []struct {
		name    string
		input   string
		kopecks int64
		amount  string
		err     error
	}{
		{
			name:    "normal",
			input:   "100.23",
			amount:  "100.23",
			kopecks: 10023,
		},
		{
			name:    "on other symbols",
			input:   "1,000,500.100",
			amount:  "1000500.10",
			kopecks: 100050010,
		},
		{
			name:    "more than 2 digits after dot",
			input:   "1.23456",
			amount:  "1.23",
			kopecks: 123,
		},
		{
			name:    "less than 2 digits after dot",
			input:   "1.2",
			amount:  "1.20",
			kopecks: 120,
		},
		{
			name:    "without dot",
			input:   "123",
			amount:  "123.00",
			kopecks: 12300,
		},
		{
			name:    "with dot and without digits after dot",
			input:   "123.",
			amount:  "123.00",
			kopecks: 12300,
		},
		{
			name:    "with dot and without digits before dot",
			input:   ".123",
			amount:  "0.12",
			kopecks: 12,
		},
		{
			name:    "rates",
			input:   "60.1662",
			amount:  "60.16",
			kopecks: 6016,
		},
		{
			name:    "amount zero",
			input:   "0",
			amount:  "0.0",
			kopecks: 0,
		},
		{
			name:    "with dot and without digits before and after dot",
			input:   ".",
			amount:  "0.00",
			kopecks: 0,
			err:     money.ErrInvalidAmount,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			kopecks, err := money.ConvertStringAmountToKopecks(tc.input)
			if tc.err != nil {
				assert.ErrorAs(t, err, &tc.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.kopecks, kopecks)

				amount := money.ConvertKopecksToAmount(kopecks)
				assert.Equal(t, tc.amount, amount)
			}
		})
	}
}
