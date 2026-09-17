package main

import (
	"flag"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

func initTracer(service string) (opentracing.Tracer, io.Closer, error) {
	tracer := opentracing.GlobalTracer()
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	return tracer, closer, nil
}

func initLogger() *zap.Logger {
	logger := zap.Must(zap.NewDevelopment())
	return logger
}

func main() {
	logger := initLogger()
	defer logger.Sync()

	tracer, closer, err := initTracer("miner")
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	span := tracer.StartSpan("miner")
	defer span.Finish()

	data := flag.String("data", "", "data to mine")
	rule := flag.Int("rule", 1234, "beginning of the hash")
	flag.Parse()

	params := url.Values{}
	params.Add("data", *data)
	params.Add("rule", strconv.Itoa(*rule))

	url := "http://localhost:9999/get_pow?" + params.Encode()
	logger.Info("Requesting the core service", zap.String("url", url))

	req, _ := http.NewRequest("GET", url, nil)

	span.SetTag("data", *data)
	span.SetTag("rule", strconv.Itoa(*rule))

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")

	err = tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	if err != nil {
		logger.Error(err.Error())
		ext.Error.Set(span, true)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error(err.Error())
		ext.Error.Set(span, true)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		ext.Error.Set(span, true)
	} else {
		logger.Info("Got the response from the core service", zap.String("proof of work", string(bodyBytes)))
	}
}
