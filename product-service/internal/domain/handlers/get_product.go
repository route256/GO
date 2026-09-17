package handlers

import (
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/domain/errors"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/domain/model"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/infra/repository/products"
)

func GetProduct(sku int64) (*model.Product, error) {

	if sku <= 0 {
		return nil, errors.ErrorInvalidSku
	}

	got := products.Get(sku)

	if got == nil {
		return nil, errors.ErrorNothingFound
	}

	return got, nil
}
