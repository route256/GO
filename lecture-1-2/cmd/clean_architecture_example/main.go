package main

import (
	"log"
	"net/http"
	"os"

	controller_http "gitlab.ozon.dev/go/classroom-8/students/lecture-1-2/internal/controller/http"
	repository "gitlab.ozon.dev/go/classroom-8/students/lecture-1-2/internal/repository/postgres"
	wms "gitlab.ozon.dev/go/classroom-8/students/lecture-1-2/internal/services/WMS"
	oms "gitlab.ozon.dev/go/classroom-8/students/lecture-1-2/internal/usecase/OMS"
	"gitlab.ozon.dev/go/classroom-8/students/lecture-1-2/pkg/middleware"
)

func main() {
	// Repository layer
	omsRepo := repository.NewOMSRepostiory( /* ... */ )

	// Other external services (adapters) layer
	wmsService := wms.NewWMSService( /* ... */ )

	// Usecase layer
	omsUsecase := oms.NewOMSUsecase(oms.Deps{
		WarehouseManagementSystem: wmsService,
		OMSRepository:             omsRepo,
	})

	// Delivery || Gateway || Transport || Controller layer
	controller := controller_http.NewController(controller_http.Usecases{
		OrderManagementSystem: omsUsecase,
	})

	// Router layer
	router := controller.NewRouter()

	// Middleware layer
	router = middleware.WithHTTPRecoverMiddleware(router)

	// Run service
	addr := os.Getenv("ADDR")
	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
