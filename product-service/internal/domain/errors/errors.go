package errors

import "errors"

var (
	ErrorInvalidCount = errors.New("count must be <= 1000 and >= 1")

	ErrorInvalidSku = errors.New("invalid SKU provided")

	ErrorNothingFound = errors.New("nothing found by the sku")
)
