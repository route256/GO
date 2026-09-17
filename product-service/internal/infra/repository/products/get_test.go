package products

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/domain/model"
)

func init() {
	InitWithDataFileName("../../../../init/data.json")
}

func TestGet(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name     string
		expected *model.Product
		givenSku int64
	}{
		{
			name:     "absent_sku",
			expected: nil,
			givenSku: int64(123),
		},
		{
			name: "present_sku",
			expected: &model.Product{
				SKU:   3596599,
				Name:  "Невербальная коммуникация. Психология и право",
				Price: 3386,
			},
			givenSku: int64(3596599),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := Get(tt.givenSku)
			require.Equalf(t, tt.expected, actual,
				fmt.Sprintf("(%s): expected %v, actual %v", tt.name, tt.expected, actual),
			)
		})
	}
}
