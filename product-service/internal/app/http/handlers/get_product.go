package handlers

import (
	"encoding/json"

	"github.com/go-openapi/runtime/middleware"
	httpModel "gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/app/http/model"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/app/http/server/operations/products"
	"gitlab.ozon.dev/go/classroom-20/experts/product-service/internal/domain/handlers"
)

func GetProduct(params products.GetProductParams, principal interface{}) middleware.Responder {

	product, err := handlers.GetProduct(params.Sku)

	if err != nil {
		return products.NewGetProductDefault(404).WithPayload(&httpModel.Error{Code: int64(1), Message: err.Error()})
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		return products.NewGetProductDefault(0).WithPayload(&httpModel.Error{Code: int64(2), Message: err.Error()})
	}

	response := new(httpModel.Product)

	err = json.Unmarshal(productJSON, response)
	if err != nil {
		return products.NewGetProductDefault(0).WithPayload(&httpModel.Error{Code: int64(3), Message: err.Error()})
	}

	return products.NewGetProductOK().WithPayload(response)
}
