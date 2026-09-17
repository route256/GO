package main

import (
	"net/http"
	"strconv"
	"workshop/core/crypto"
	"workshop/core/metrics"
	"workshop/core/trace"

	"workshop/core/log"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

const serviceName = "core"

func main() {

	logger, err := log.InitLogger(serviceName)
	if err != nil {
		panic(err)
	}
	logger.Info("The service has started")

	tracer, closer, err := trace.InitTracer(serviceName)
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	http.HandleFunc("/get_pow", func(w http.ResponseWriter, r *http.Request) {

		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		span := tracer.StartSpan("getting pow", ext.RPCServerOption(spanCtx))
		defer span.Finish()
		defer metrics.RequestsTotal.WithLabelValues("200").Inc()

		data := r.URL.Query().Get("data")
		rule := r.URL.Query().Get("rule")
		logger.Info("Received some data", zap.String("data", data), zap.String("rule", rule))

		pow, hash, err := crypto.Mine(data, rule)
		if err != nil {
			metrics.RequestsTotal.WithLabelValues("500").Inc()
			logger.Error(err.Error())
			ext.Error.Set(span, true)
			return
		} else {
			logger.Info("Processed the request", zap.Int64("proof of work", pow), zap.String("hash", hash))
		}

		span.SetTag("pow", pow)
		span.SetTag("hash", hash)

		response := []byte(strconv.FormatInt(pow, 10))
		_, err = w.Write(response)
		if err != nil {
			metrics.RequestsTotal.WithLabelValues("500").Inc()
			logger.Error(err.Error())
			ext.Error.Set(span, true)
			return
		}
	})
	http.Handle("/metrics", promhttp.Handler())

	err = http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
