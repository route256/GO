package controller_http

import (
	"net/http"
)

// NewRouter - returns http.Handler
func (c *Controller) NewRouter() http.Handler {
	// Router layer
	mux := http.NewServeMux()

	// Note: You can add here custom middleware too
	mux.HandleFunc("/v1/order/create", c.CreateOrderHandler)

	return mux
}
