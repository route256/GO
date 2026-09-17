package products

import "gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/domain/model"

// Get -
func Get(key int64) *model.Product {

	got, found := productsMap[key]
	if !found {
		return nil
	}

	copy := new(model.Product)

	*copy = *got

	return copy
}
