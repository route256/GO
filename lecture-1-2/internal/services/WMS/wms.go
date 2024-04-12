package wms

import (
	"context"
	"errors"

	"gitlab.ozon.dev/go/classroom-8/students/lecture-1-2/internal/models"
	oms "gitlab.ozon.dev/go/classroom-8/students/lecture-1-2/internal/usecase/OMS"
)

type wmsService struct {
	/*
		HTTP, gRPC, XML, ... client
	*/
}

// Check that we implemet contract for usecase
var _ oms.WarehouseManagementSystem = (*wmsService)(nil)

// NewWMSService - returns WMS service adapter
func NewWMSService( /* ... */ ) *wmsService {
	return &wmsService{
		/* ... */
	}
}

func (r *wmsService) ReserveStocks(ctx context.Context, userID models.UserID, items []models.ItemOrderInfo) error {
	return errors.New("unimplemented")
}
