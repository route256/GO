package handlers

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/domain/errors"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/domain/model"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/infra/repository/products"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/infra/repository/skus"
)

func init() {
	skus.InitWithDataFileName("../../../init/skus.json")
	products.InitWithDataFileName("../../../init/data.json")
}

func TestGetProduct(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name            string
		expectedErr     error
		expectedProduct *model.Product
		given           int64
	}{
		{
			name:        "invalid_sku",
			expectedErr: errors.ErrorInvalidSku,
			given:       -1,
		},
		{
			name:        "absent_sku",
			expectedErr: errors.ErrorNothingFound,
			given:       123,
		},
		{
			name:        "present_sku",
			expectedErr: nil,
			given:       1625903,
			expectedProduct: &model.Product{
				SKU:   1625903,
				Price: 1423,
				Name:  "Сказки звездного неба",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := GetProduct(tt.given)

			if tt.expectedErr == nil {
				require.NoError(t, err)
			} else {
				require.Equal(t, tt.expectedErr, err)
			}

			if tt.expectedProduct != nil {
				require.Equal(t, actual, tt.expectedProduct)
			} else {
				require.Nil(t, actual)
			}
		})
	}
}
