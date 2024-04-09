package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	InFlightRequests = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "ozon",
		Subsystem: "http",
		Name:      "in_flight_requests_total",
	})
	SummaryResponseTime = promauto.NewSummary(prometheus.SummaryOpts{
		Namespace: "ozon",
		Subsystem: "http",
		Name:      "summary_response_time_seconds",
		Objectives: map[float64]float64{
			0.5:  0.1,
			0.9:  0.01,
			0.99: 0.001,
		},
	})
	HistogramResponseTime = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "ozon",
			Subsystem: "http",
			Name:      "histogram_response_time_seconds",
			Buckets:   []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2},
			// Buckets: prometheus.ExponentialBucketsRange(0.0001, 2, 16),
		},
		[]string{"code"},
	)
)

func MetricsMiddleware(next http.Handler) http.Handler {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper := NewResponseWrapper(w)

		startTime := time.Now()
		next.ServeHTTP(wrapper, r)
		duration := time.Since(startTime)

		SummaryResponseTime.Observe(duration.Seconds())
		HistogramResponseTime.
			WithLabelValues(http.StatusText(wrapper.statusCode)).
			Observe(duration.Seconds())
	})
	wrappedHandler := promhttp.InstrumentHandlerInFlight(InFlightRequests, handler)
	return wrappedHandler
}
