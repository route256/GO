package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var (
	port        = flag.Int("port", 8080, "the port to listen")
	develMode   = flag.Bool("devel", false, "development mode")
	serviceName = flag.String("service", "fibonacci", "the name of our service")
)

func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		numStr := r.URL.Query().Get("n")
		num, err := strconv.Atoi(numStr)
		if err != nil || num < 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, getNumber(ctx, num))
	})
}

func main() {
	flag.Parse()

	logger := initLogger()
	initTracing(logger)

	handler := Handler()
	handler = LoggingMiddleware(logger, handler)
	handler = MetricsMiddleware(handler)
	handler = TracingMiddleware(handler)

	multiHandler := MultiHandler()
	multiHandler = LoggingMiddleware(logger, multiHandler)
	multiHandler = TracingMiddleware(multiHandler)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/fibonacci", handler)
	http.Handle("/multi", multiHandler)

	// sugaredLogger := logger.Sugar()
	// sugaredLogger.Infow("starting http server", "port", *port)

	logger.Info("starting http server", zap.Int("port", *port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		logger.Fatal("error starting http server", zap.Error(err))
	}
}
