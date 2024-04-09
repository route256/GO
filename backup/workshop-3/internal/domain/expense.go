package domain

import "time"

// Я вынес доменные сущности из database чтобы не тянуть везде зависиомсть
// от пакета database и воспользоваться утинной типизацией в командах для бота
type Expense struct {
	Title  string
	Date   time.Time
	Amount int64 // в копейках
}
