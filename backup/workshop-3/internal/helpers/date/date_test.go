package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDate(t *testing.T) {
	// тест простой и может излишний, но тест написан для того чтобы никто случайно не поменял логику самой функции GetDate
	t.Run("method should leave only the day month and year", func(t *testing.T) {
		now := time.Now()
		date := GetDate(now)
		expectedDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

		assert.EqualValues(t, expectedDate, date)
	})

}
