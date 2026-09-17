package usecase

import "errors"

var (
	ErrReserveStocks = errors.New("can't reserve stocks")
	ErrCreateOrder   = errors.New("can't create order")
)
