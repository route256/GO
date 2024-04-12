package models

import (
	"time"

	"github.com/google/uuid"
)

// Order - заказ
type Order struct {
	UUID              uuid.UUID       // UUID заказа
	UserID            UserID          // ID пользователя (чей заказ)
	Items             []ItemOrderInfo // Информация о составе заказа
	DeliveryOrderInfo                 // Информация о доставке
	/* ... */
}

// DeliveryOrderInfo - информация о доставке заказа
type DeliveryOrderInfo struct {
	DeliveryVariantID DeliveryVariantID
	DeliveryDate      time.Time
}

// ItemOrderInfo - информация о составе заказа
type ItemOrderInfo struct {
	SKU         SKU         // SKU
	Quantity    uint16      // количество SKU
	WarehouseID WarehouseID // с какого склада будет браться сток
}
