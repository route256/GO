package oms

import (
	"context"

	"github.com/google/uuid"
	"gitlab.ozon.dev/go/classroom-8/students/lecture-1-2/internal/models"
	"gitlab.ozon.dev/go/classroom-8/students/lecture-1-2/internal/usecase"
)

// Объявляем интерфейсы зависимостей в месте использования!
// (Задаем контракт поведения для адаптеров)
type (
	// WarehouseManagementSystem - то что отвечает за стоки товаров
	WarehouseManagementSystem interface {
		// ReserveStocks - резервация стоков на складах
		ReserveStocks(ctx context.Context, userID models.UserID, items []models.ItemOrderInfo) error
	}

	// OMSRepository - репозиторий сервиса OMS
	OMSRepository interface {
		// CreateOrder - создание записи заказа в БД
		CreateOrder(ctx context.Context, order models.Order) error
	}
)

// Deps - зависимости нашего usecase
type Deps struct {
	WarehouseManagementSystem
	OMSRepository
}

type omsUsecase struct {
	Deps
}

// check that we implement usecase contarct correctly
var _ usecase.OrderManagementSystem = (*omsUsecase)(nil)

// NewOMSUsecase - возвращаем реализацию usecase.OrderManagementSystem
func NewOMSUsecase(d Deps) *omsUsecase {
	return &omsUsecase{
		Deps: d,
	}
}

// CreateOrder - создание заказа
func (oms *omsUsecase) CreateOrder(ctx context.Context, userID models.UserID, info usecase.CreateOrderInfo) (models.Order, error) {
	// Резервируем стоки на складах
	if err := oms.WarehouseManagementSystem.ReserveStocks(ctx, userID, info.Items); err != nil {
		return models.Order{}, usecase.ErrReserveStocks
	}

	// Формируем запись о заказе
	var (
		orderUUID = uuid.New()
		order     = models.Order{
			UUID:              orderUUID,
			UserID:            userID,
			Items:             info.Items,
			DeliveryOrderInfo: info.DeliveryOrderInfo,
		}
	)

	// Создаем заказ в БД
	if err := oms.OMSRepository.CreateOrder(ctx, order); err != nil {
		return models.Order{}, usecase.ErrCreateOrder
	}

	return order, nil
}
