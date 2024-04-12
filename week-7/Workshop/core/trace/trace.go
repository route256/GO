package trace

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitTracer(serviceName string) (opentracing.Tracer, io.Closer, error) {
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	return tracer, closer, err
}
