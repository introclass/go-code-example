package main

import (
	"go-code-example/opentracingusage"
	"log"
	"net/http"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func server() {
	tracer := opentracing.GlobalTracer()
	http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		spanctx, _ := tracer.Extract(opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		serverspan := tracer.StartSpan("server", ext.RPCServerOption(spanctx))
		defer serverspan.Finish()
	})
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func client() {
	url := "http://localhost:8082/publish"
	req, _ := http.NewRequest("GET", url, nil)

	tracer := opentracing.GlobalTracer()

	clientSpan := tracer.StartSpan("client")
	defer clientSpan.Finish()

	ext.SpanKindRPCClient.Set(clientSpan)
	ext.HTTPUrl.Set(clientSpan, url)
	ext.HTTPMethod.Set(clientSpan, "GET")

	tracer.Inject(clientSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	http.DefaultClient.Do(req)
}

func main() {
	tracer, closer, err := opentracingusage.InitJaeger("01-httptrace", "127.0.0.1:6831")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	go server()

	time.Sleep(1 * time.Second)
	client()
}
