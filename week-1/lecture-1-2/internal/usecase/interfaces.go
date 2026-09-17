package usecase

import (
	"context"

	"gitlab.ozon.dev/go/classroom-8/students/lecture-1-2/internal/models"
)

// interfaces.go: Декларируем бизнес функциональность

type OrderManagementSystem interface {
	CreateOrder(ctx context.Context, userID models.UserID, info CreateOrderInfo) (models.Order, error)
	/* ... */
}
