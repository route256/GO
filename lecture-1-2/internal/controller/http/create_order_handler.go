package controller_http

import (
	"encoding/json"
	"net/http"
	"time"

	"gitlab.ozon.dev/go/classroom-8/students/lecture-1-2/internal/models"
	"gitlab.ozon.dev/go/classroom-8/students/lecture-1-2/internal/usecase"
)

type CreateOrderRequest struct {
	UserID int64 `json:"user_id"` // ID пользователя
	Items  []struct {
		ID          int64  `json:"id"`           // ID товара
		Quantity    uint16 `json:"quantity"`     // Количество данного товара
		Price       uint32 `json:"price"`        // Цена одного товара
		WarehouseID int64  `json:"warehouse_id"` // ID склада с которого поедет товар
	} `json:"items"` // Товары в корзине к оплате
	DeliveryVariantID int64     `json:"delivery_variant_id"` // ID способа доставки
	DelieveryDate     time.Time `json:"delivery_date"`       // Срок доставки
}

type CreateOrderResponse struct {
	OrderUUID string `json:"order_uuid"` // UUID созданного заказа
}

func (c *Controller) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	// 0. Decode request
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 1. Validation
	if err := validateCreateOrderRequest(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Transform delivery layer models to Domain/Usecase models
	orderInfo := extractCreateOrderInfoFromCreateOrderRequest(&req)

	// 3. Call usecases
	order, err := c.OrderManagementSystem.CreateOrder(ctx, models.UserID(req.UserID), orderInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Prepare answer
	resp := CreateOrderResponse{
		OrderUUID: order.UUID.String(),
	}

	// 5. Encode answer & send response
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func validateCreateOrderRequest(req *CreateOrderRequest) error {
	/* your validation logic here */
	return nil
}

func extractCreateOrderInfoFromCreateOrderRequest(req *CreateOrderRequest) usecase.CreateOrderInfo {
	/* your mapping logic here */

	info := usecase.CreateOrderInfo{
		/* ... */
	}

	return info
}
