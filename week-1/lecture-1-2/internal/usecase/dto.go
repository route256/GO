package usecase

import (
	"gitlab.ozon.dev/go/classroom-8/students/lecture-1-2/internal/models"
)

// CreateOrderInputInfo - DTO заказа (для создания заказа)
type CreateOrderInfo struct {
	Items                    []models.ItemOrderInfo // Товары в заказе
	models.DeliveryOrderInfo                        // Информация о доставке
}
