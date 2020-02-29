package main

import (
	"io"
	"log"

	opentracing "github.com/opentracing/opentracing-go"
	oplog "github.com/opentracing/opentracing-go/log"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

func initJaeger(service, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := jaegercfg.Configuration{
		ServiceName: service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: addr,
		},
	}
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	return cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
}

//docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest
func main() {
	tracer, closer, err := initJaeger("00-quickstart", "127.0.0.1:6831")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	rootspan := opentracing.StartSpan("rootcall")
	defer rootspan.Finish()
	rootspan.SetTag("rootcall", "rootcall")

	subspan := opentracing.StartSpan("subcall",
		opentracing.ChildOf(rootspan.Context()))
	defer subspan.Finish()
	subspan.SetTag("subcall", "subcall")
	subspan.LogFields(oplog.String("event", "testlog"))
}
