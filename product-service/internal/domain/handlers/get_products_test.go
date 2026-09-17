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

func TestGetProducts(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name             string
		expectedErr      error
		expectedProducts []*model.Product
		givenLimit       int32
		givenSkip        int64
	}{
		{
			name:        "invalid_limit_1",
			givenLimit:  -1,
			expectedErr: errors.ErrorInvalidCount,
		},
		{
			name:        "invalid_limit_2",
			givenLimit:  2000,
			expectedErr: errors.ErrorInvalidCount,
		},
		{
			name:        "invalid_skip",
			givenLimit:  2,
			givenSkip:   -2,
			expectedErr: errors.ErrorInvalidSku,
		},
		{
			name:       "correct",
			givenLimit: 2,
			givenSkip:  0,
			expectedProducts: []*model.Product{
				{
					Name:  "Теория нравственных чувств | Смит Адам",
					Price: 3379,
					SKU:   1076963,
				},
				{
					Name:  "Кулинар Гуров",
					Price: 2931,
					SKU:   1148162,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := GetProducts(tt.givenSkip, tt.givenLimit)

			if tt.expectedErr == nil {
				require.NoError(t, err)
			} else {
				require.Equal(t, tt.expectedErr, err)
			}

			if tt.expectedProducts != nil {
				require.Equal(t, actual, tt.expectedProducts)
			} else {
				require.Nil(t, actual)
			}
		})
	}
}
