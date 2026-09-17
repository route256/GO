package handlers

import (
	"encoding/json"

	"github.com/go-openapi/runtime/middleware"
	httpModel "gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/app/http/model"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/app/http/server/operations/products"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/domain/handlers"
)

func GetProducts(params products.GetProductsParams, principal interface{}) middleware.Responder {

	domainProducts, err := handlers.GetProducts(*params.StartAfterSku, *params.Count)

	if err != nil {
		return products.NewGetProductsDefault(0).WithPayload(&httpModel.Error{Code: 1, Message: err.Error()})
	}

	answer := make([]*httpModel.Product, 0, len(domainProducts))

	for _, product := range domainProducts {

		productJSON, err := json.Marshal(product)
		if err != nil {
			return products.NewGetProductsDefault(0).WithPayload(&httpModel.Error{Code: 2, Message: err.Error()})
		}

		httpProduct := new(httpModel.Product)

		err = json.Unmarshal(productJSON, &httpProduct)
		if err != nil {
			return products.NewGetProductsDefault(0).WithPayload(&httpModel.Error{Code: 3, Message: err.Error()})
		}

		answer = append(answer, httpProduct)
	}

	return products.NewGetProductsOK().WithPayload(answer)

}
