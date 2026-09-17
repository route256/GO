package handlers

import (
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/domain/errors"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/domain/model"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/infra/repository/products"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/infra/repository/skus"
)

const (
	ListSkuThrh = 1000
)

func GetProducts(skip int64, limit int32) ([]*model.Product, error) {

	if limit > ListSkuThrh || limit <= 0 {
		return nil, errors.ErrorInvalidCount
	}

	if skip < 0 {
		return nil, errors.ErrorInvalidSku
	}

	skus := skus.List(skip, limit)
	answer := make([]*model.Product, 0, len(skus))
	for _, sku := range skus {
		model := products.Get(sku)
		if model != nil {
			answer = append(answer, model)
		}
	}

	return answer, nil
}
